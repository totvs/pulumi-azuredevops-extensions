package provider

import (
	"fmt"
	"strings"

	"encoding/base64"

	resty "github.com/go-resty/resty/v2"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AzureDevopsBuildFolderPermissionsResource struct {
	config AzureDevopsConfig
}

type AzureDevopsBuildPermissions struct {
	Allow int64
	Deny  int64
}

type ExtendedInfo struct {
	EffectiveAllow int64 `json:"effectiveAllow"`
	EffectiveDeny  int64 `json:"effectiveDeny"`
	InheritedAllow int64 `json:"inheritedAllow"`
	InheritedDeny  int64 `json:"inheritedDeny"`
}

type AccessControlEntry struct {
	Descriptor   string       `json:"descriptor"`
	Allow        int64        `json:"allow"`
	Deny         int64        `json:"deny"`
	ExtendedInfo ExtendedInfo `json:"extendedInfo"`
}

type AccessControlEntries struct {
	Token                string               `json:"token"`
	Merge                bool                 `json:"merge"`
	AccessControlEntries []AccessControlEntry `json:"accessControlEntries"`
}

const (
	BUILD_SECURITY_NAMESPACE_ID       = "33344d9c-fc72-4d6f-aba5-fa317101a7e9"
	MICROSOFT_TEAMFOUNDATION_IDENTITY = "Microsoft.TeamFoundation.Identity"
)

type AzureDevopsBuildFolderPermissionsInput struct {
	ProjectId   string               `json:"projectId"`
	Principal   string               `json:"principal"`
	Path        string               `json:"path"`
	Permissions resource.PropertyMap `json:"permissions"`
	Replace     bool                 `json:"replace"`
}

func (c *AzureDevopsBuildFolderPermissionsResource) Name() string {
	return "azuredevops-extensions:index:BuildFolderPermissions"
}

func (c *AzureDevopsBuildFolderPermissionsResource) Configure(config AzureDevopsConfig) {
	c.config = config
}

func (c *AzureDevopsBuildFolderPermissionsResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: false})
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

	var replaces []string
	if diffsInput.Changed("path") {
		replaces = append(replaces, "path")
	}
	if diffsInput.Changed("permissions") {
		replaces = append(replaces, "permissions")
	}
	if diffsInput.Changed("principal") {
		replaces = append(replaces, "principal")
	}
	if diffsInput.Changed("projectId") {
		replaces = append(replaces, "projectId")
	}
	if diffsInput.Changed("replace") {
		replaces = append(replaces, "replace")
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_SOME,
		Replaces:            replaces,
		Stables:             []string{},
		DeleteBeforeReplace: true,
	}, nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	buildFolderPermissionsInputs := c.ToAzureDevopsBuildFolderPermissionsInput(inputs)
	buildFolderPermissions, err := c.createBuildFolderPermissions(buildFolderPermissionsInputs)
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
		Id:         buildFolderPermissions,
		Properties: outputProperties,
	}, nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputsBuildFolderPermissions := c.ToAzureDevopsBuildFolderPermissionsInput(inputs["__inputs"].ObjectValue())

	return &pbempty.Empty{}, c.removeBuildFolderPermissions(inputsBuildFolderPermissions)
}

func (c *AzureDevopsBuildFolderPermissionsResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, fmt.Errorf("Update is not supported")
}

func (k *AzureDevopsBuildFolderPermissionsResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "read is not yet implemented for "+k.Name())
}

func (r *AzureDevopsBuildFolderPermissionsResource) ToAzureDevopsBuildFolderPermissionsInput(inputMap resource.PropertyMap) AzureDevopsBuildFolderPermissionsInput {
	input := AzureDevopsBuildFolderPermissionsInput{}

	if inputMap["projectId"].HasValue() && inputMap["projectId"].IsString() {
		input.ProjectId = inputMap["projectId"].StringValue()
	}

	if inputMap["principal"].HasValue() && inputMap["principal"].IsString() {
		input.Principal = inputMap["principal"].StringValue()
	}

	if inputMap["path"].HasValue() && inputMap["path"].IsString() {
		input.Path = inputMap["path"].StringValue()
	}

	if inputMap["permissions"].HasValue() && inputMap["permissions"].IsObject() {
		input.Permissions = inputMap["permissions"].ObjectValue()
	}

	if inputMap["replace"].HasValue() && inputMap["replace"].IsBool() {
		input.Replace = inputMap["replace"].BoolValue()
	}

	return input
}

func (c *AzureDevopsBuildFolderPermissionsResource) createBuildFolderPermissions(buildFolderPermissionInput AzureDevopsBuildFolderPermissionsInput) (string, error) {
	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return "", err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return "", err
	}

	accessControlEntriesRequestBody, err := c.createAccessControlEntriesRequestBody(
		buildFolderPermissionInput.Principal,
		buildFolderPermissionInput.ProjectId,
		buildFolderPermissionInput.Path,
		buildFolderPermissionInput.Permissions,
		buildFolderPermissionInput.Replace,
	)
	if err != nil {
		return "", err
	}

	// TODO: implement aad user and others...
	if !strings.HasPrefix(accessControlEntriesRequestBody.AccessControlEntries[0].Descriptor, "Microsoft.TeamFoundation.Identity") {
		return "", fmt.Errorf("access control not supported: %s", accessControlEntriesRequestBody.AccessControlEntries[0].Descriptor)
	}

	client := resty.New()
	url := fmt.Sprintf("%s/_apis/AccessControlEntries/{securityNamespace}", *urlOrg)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetPathParam("securityNamespace", BUILD_SECURITY_NAMESPACE_ID).
		SetQueryString("api-version=6.0").
		SetBody(accessControlEntriesRequestBody).
		Post(url)

	if err != nil || resp.StatusCode() != 200 {
		message := ""
		azError, err := MarshalAzureDevopsError(resp.Body())
		if err == nil {
			message = azError.Message
		}
		return "", fmt.Errorf(
			"error creating build folder permission [%s, %s, %s, %s, %s]': %s",
			*urlOrg,
			BUILD_SECURITY_NAMESPACE_ID,
			buildFolderPermissionInput.ProjectId,
			buildFolderPermissionInput.Path,
			resp.Status(),
			message)
	}

	return c.createBuildFolderPermissionsId(accessControlEntriesRequestBody), err
}

func (c *AzureDevopsBuildFolderPermissionsResource) createAccessControlEntriesRequestBody(principal string, projectId string, path string, list resource.PropertyMap, replace bool) (AccessControlEntries, error) {
	if !strings.HasPrefix(principal, "vssgp.") {
		return AccessControlEntries{}, fmt.Errorf("principal not supported: %s", principal)
	}

	descriptor, err := c.createDescriptor(principal)
	if err != nil {
		return AccessControlEntries{}, err
	}

	buildPermissions := c.getBuildPermissions(list)

	return AccessControlEntries{
		Token: c.getAzureDevopsPermissionsToken(projectId, path),
		Merge: replace,
		AccessControlEntries: []AccessControlEntry{
			{
				Descriptor: descriptor,
				Allow:      buildPermissions.Allow,
				Deny:       buildPermissions.Deny,
				ExtendedInfo: ExtendedInfo{
					EffectiveAllow: buildPermissions.Allow,
					EffectiveDeny:  buildPermissions.Deny,
					InheritedAllow: buildPermissions.Allow,
					InheritedDeny:  buildPermissions.Deny,
				},
			},
		},
	}, nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) createBuildFolderPermissionsId(accessControlEntries AccessControlEntries) string {
	return fmt.Sprintf(
		"%s\\%d\\%d\\%s",
		accessControlEntries.Token,
		accessControlEntries.AccessControlEntries[0].Allow,
		accessControlEntries.AccessControlEntries[0].Deny,
		accessControlEntries.AccessControlEntries[0].Descriptor,
	)
}

func (c *AzureDevopsBuildFolderPermissionsResource) removeBuildFolderPermissions(buildFolderPermissionInput AzureDevopsBuildFolderPermissionsInput) error {
	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return err
	}

	accessControlEntriesRequestBody, err := c.createAccessControlEntriesRequestBody(
		buildFolderPermissionInput.Principal,
		buildFolderPermissionInput.ProjectId,
		buildFolderPermissionInput.Path,
		buildFolderPermissionInput.Permissions,
		buildFolderPermissionInput.Replace,
	)
	if err != nil {
		return err
	}

	client := resty.New()
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetPathParam("securityNamespace", BUILD_SECURITY_NAMESPACE_ID).
		SetQueryParam("token", accessControlEntriesRequestBody.Token).
		SetQueryParam("descriptors", accessControlEntriesRequestBody.AccessControlEntries[0].Descriptor).
		SetHeader("Accept", "application/json;api-version=5.0-preview.1;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true").
		Delete(fmt.Sprintf("%s/_apis/AccessControlEntries/{securityNamespace}", *urlOrg))

	if err != nil || resp.StatusCode() != 200 {
		message := ""
		azError, err := MarshalAzureDevopsError(resp.Body())
		if err == nil {
			message = azError.Message
		}
		return fmt.Errorf(
			"error deleting build folder permission [%s, %s, %s, %s, %s]': %s",
			*urlOrg,
			BUILD_SECURITY_NAMESPACE_ID,
			buildFolderPermissionInput.ProjectId,
			buildFolderPermissionInput.Path,
			resp.Status(),
			message)

	}

	return nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) createDescriptor(principal string) (string, error) {
	tokens := strings.Split(principal, ".")
	if len(tokens) != 2 {
		return "", fmt.Errorf("invalid principal format: %s", principal)
	}

	decoded, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(tokens[1] + "=")
		if err != nil {
			return "", fmt.Errorf("invalid principal format: %s (%s)", principal, err)
		}
	}
	descriptor := fmt.Sprintf("%s;%s", MICROSOFT_TEAMFOUNDATION_IDENTITY, string(decoded))

	return descriptor, nil
}

func (c *AzureDevopsBuildFolderPermissionsResource) getAzureDevopsPermissionsToken(projectId string, path string) string {
	pathNormalized := strings.Replace(path, "\\", "/", -1)
	return fmt.Sprintf("%s%s", projectId, pathNormalized)
}

func (c *AzureDevopsBuildFolderPermissionsResource) getBuildPermissions(list resource.PropertyMap) AzureDevopsBuildPermissions {
	permissions := AzureDevopsBuildPermissions{}

	for k, v := range list {
		if v.String() == "{Allow}" {
			permissions.Allow |= c.getBuildPermissionBitMask(string(k))
		} else if v.String() == "{Deny}" {
			permissions.Deny |= c.getBuildPermissionBitMask(string(k))
		}
	}

	return permissions
}

func (c *AzureDevopsBuildFolderPermissionsResource) getBuildPermissionBitMask(permission string) int64 {
	switch permission {
	case "ViewBuilds":
		return 1
	case "EditBuildQuality":
		return 2
	case "RetainIndefinitely":
		return 4
	case "DeleteBuilds":
		return 8
	case "ManageBuildQualities":
		return 16
	case "DestroyBuilds":
		return 32
	case "UpdateBuildInformation":
		return 64
	case "QueueBuilds":
		return 128
	case "ManageBuildQueue":
		return 256
	case "StopBuilds":
		return 512
	case "ViewBuildDefinition":
		return 1024
	case "EditBuildDefinition":
		return 2048
	case "DeleteBuildDefinition":
		return 4096
	case "OverrideBuildCheckInValidation":
		return 8192
	case "AdministerBuildPermissions":
		return 16384
	}

	return 0
}
