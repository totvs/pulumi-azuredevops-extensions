package provider

import (
	"fmt"
	"strings"

	resty "github.com/go-resty/resty/v2"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AzureDevopsBuildFolderResource struct {
	config AzureDevopsConfig
}

type AzureDevopsBuildFolderInputId struct {
	ProjectId string `json:"projectId"`
	Path      string `json:"path"`
}

func (c *AzureDevopsBuildFolderResource) Name() string {
	return "azuredevops-extensions:index:BuildFolder"
}

func (c *AzureDevopsBuildFolderResource) Configure(config AzureDevopsConfig) {
	c.config = config
}

func (c *AzureDevopsBuildFolderResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: false, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: false})
	if err != nil {
		return nil, err
	}

	diffsInput := olds["__inputs"].ObjectValue().Diff(news)
	if diffsInput == nil {
		return &pulumirpc.DiffResponse{
			Changes:             pulumirpc.DiffResponse_DIFF_NONE,
			DeleteBeforeReplace: false,
		}, nil
	}

	var diffs []string
	if diffsInput.Changed("path") {
		diffs = append(diffs, "path")
	}

	var replaces []string
	if diffsInput.Changed("projectId") {
		replaces = append(replaces, "projectId")
	}

	if len(replaces) > 0 {
		replaces = append(replaces, diffs...)
		diffs = nil
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_SOME,
		Replaces:            replaces,
		Diffs:               diffs,
		Stables:             []string{},
		DeleteBeforeReplace: true,
	}, nil
}

func (c *AzureDevopsBuildFolderResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputsBuildFolderId := c.ToAzureDevopsBuildFolderInputId(inputs)
	buildFolderId, err := c.createBuildFolder(inputsBuildFolderId)
	if err != nil {
		return nil, err
	}

	outputStore := resource.PropertyMap{}
	outputStore["__inputs"] = resource.NewObjectProperty(inputs)

	outputProperties, err := plugin.MarshalProperties(
		outputStore,
		plugin.MarshalOptions{},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         *buildFolderId,
		Properties: outputProperties,
	}, nil
}

func (c *AzureDevopsBuildFolderResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	values := inputs["__inputs"].ObjectValue()
	inputsBuildFolderId := c.ToAzureDevopsBuildFolderInputId(values)

	return &pbempty.Empty{}, c.removeBuildFolder(inputsBuildFolderId)
}

func (c *AzureDevopsBuildFolderResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (c *AzureDevopsBuildFolderResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	inputOlds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputNews, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputOldsBuildFolder := c.ToAzureDevopsBuildFolderInputId(inputOlds["__inputs"].ObjectValue())
	inputNewsBuildFolder := c.ToAzureDevopsBuildFolderInputId(inputNews)
	_, err = c.updateBuildFolder(inputOldsBuildFolder, inputNewsBuildFolder)
	if err != nil {
		return nil, err
	}

	outputStore := resource.PropertyMap{}
	outputStore["__inputs"] = resource.NewObjectProperty(inputNews)

	outputProperties, err := plugin.MarshalProperties(
		outputStore,
		plugin.MarshalOptions{},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

func (k *AzureDevopsBuildFolderResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "read is not yet implemented for "+k.Name())
}

func (r *AzureDevopsBuildFolderResource) ToAzureDevopsBuildFolderInputId(inputMap resource.PropertyMap) AzureDevopsBuildFolderInputId {
	input := AzureDevopsBuildFolderInputId{}

	if inputMap["projectId"].HasValue() && inputMap["projectId"].IsString() {
		input.ProjectId = inputMap["projectId"].StringValue()
	}

	if inputMap["path"].HasValue() && inputMap["path"].IsString() {
		input.Path = inputMap["path"].StringValue()
	}

	return input
}

func (c *AzureDevopsBuildFolderResource) createBuildFolder(buildFolderId AzureDevopsBuildFolderInputId) (*string, error) {
	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return nil, err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	url := fmt.Sprintf(
		"%s/{projectId}/_apis/build/folders",
		*urlOrg,
	)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetPathParam("projectId", buildFolderId.ProjectId).
		SetQueryParam("path", buildFolderId.Path).
		SetQueryString("api-version=6.0-preview.2").
		SetHeader("Content-Type", "application/json").
		SetBody("{}").
		Put(url)

	if err != nil || resp.StatusCode() != 200 {
		message := ""
		azError, err := MarshalAzureDevopsError(resp.Body())
		if err == nil {
			message = azError.Message
		}
		return nil, fmt.Errorf(
			"error creating build folder [%s, %s, %s, %s]': %s",
			*urlOrg,
			buildFolderId.ProjectId,
			buildFolderId.Path,
			resp.Status(),
			message)
	}

	id := c.createBuildFolderId(buildFolderId)

	return &id, err
}

func (c *AzureDevopsBuildFolderResource) updateBuildFolder(
	oldBuildFolderId AzureDevopsBuildFolderInputId,
	newBuildFolderId AzureDevopsBuildFolderInputId) (*string, error) {

	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return nil, err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return nil, err
	}

	newPath := fmt.Sprintf(`{
		"path": "%s"
	}`, transformPath(newBuildFolderId.Path))

	client := resty.New()
	url := fmt.Sprintf(
		"%s/{projectId}/_apis/build/folders",
		*urlOrg,
	)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.0-preview.2").
		SetPathParam("projectId", oldBuildFolderId.ProjectId).
		SetQueryParam("path", oldBuildFolderId.Path).
		SetHeader("Content-Type", "application/json").
		SetBody(newPath).
		Post(url)

	if err != nil || resp.StatusCode() != 200 {
		message := ""
		azError, err := MarshalAzureDevopsError(resp.Body())
		if err == nil {
			message = azError.Message
		}
		return nil, fmt.Errorf(
			"error creating build folder [%s, %s, %s, %s, %s]': %s",
			*urlOrg,
			oldBuildFolderId.ProjectId,
			oldBuildFolderId.Path,
			newBuildFolderId.Path,
			resp.Status(),
			message)
	}

	id := c.createBuildFolderId(newBuildFolderId)

	return &id, err
}

func (c *AzureDevopsBuildFolderResource) createBuildFolderId(input AzureDevopsBuildFolderInputId) string {
	return fmt.Sprintf(
		"%s\\%s",
		input.ProjectId,
		input.Path,
	)
}

func (c *AzureDevopsBuildFolderResource) removeBuildFolder(buildFolderId AzureDevopsBuildFolderInputId) error {
	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return err
	}

	client := resty.New()
	url := fmt.Sprintf(
		"%s/{projectId}/_apis/build/folders",
		*urlOrg,
	)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetPathParam("projectId", buildFolderId.ProjectId).
		SetQueryParam("path", buildFolderId.Path).
		SetQueryString("api-version=6.0-preview.2").
		Delete(url)

	if err != nil || resp.StatusCode() != 204 {
		return fmt.Errorf(
			"error creating build folder [%s, %s, %s, %s]': %s",
			*urlOrg,
			buildFolderId.ProjectId,
			buildFolderId.Path,
			resp.Status(),
			err)
	}

	return err
}

func transformPath(path string) string {
	return strings.Replace(path, "/", "\\\\", -1)
}
