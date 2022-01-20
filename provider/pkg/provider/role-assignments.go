package provider

import (
	"encoding/json"
	"fmt"
	"log"

	resty "github.com/go-resty/resty/v2"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AzureDevopsRoleAssignmentResource struct {
	config AzureDevopsConfig
}

type ScopeNameInput string

type RoleNameInput string

type RoleAssignment struct {
	RoleName string `json:"roleName"`
	UserId   string `json:"userId"`
}

type AzureDevopsRoleAssignmentInput struct {
	Id       AzureDevopsRoleAssignmentId `json:"id"`
	RoleName RoleNameInput               `json:"roleName"`
}

type AzureDevopsRoleAssignmentId struct {
	ResourceId string         `json:"resourceId"`
	IdentityId string         `json:"identityId"`
	ScopeName  ScopeNameInput `json:"scopeName"`
	UserId     string         `json:"userId"`
}

const (
	VariableGroup ScopeNameInput = "VariableGroup"
)

const (
	Administrator RoleNameInput = "Administrator"
	User          RoleNameInput = "User"
	Reader        RoleNameInput = "Reader"
)

func (a *ScopeNameInput) GetScopeId() string {
	switch *a {
	case VariableGroup:
		return "distributedtask.variablegroup"
	}

	return ""
}

func (c *AzureDevopsRoleAssignmentResource) Name() string {
	return "azuredevops-extensions:index:RoleAssignment"
}

func (c *AzureDevopsRoleAssignmentResource) Configure(config AzureDevopsConfig) {
	c.config = config
}

func (c *AzureDevopsRoleAssignmentResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: false, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: false})
	if err != nil {
		return nil, err
	}

	diffs := olds["__inputs"].ObjectValue().Diff(news)
	if diffs == nil {
		return &pulumirpc.DiffResponse{
			Changes:             pulumirpc.DiffResponse_DIFF_NONE,
			Replaces:            []string{},
			Stables:             []string{},
			DeleteBeforeReplace: false,
		}, nil
	}

	var replaces []string
	if diffs.Changed("roleName") {
		replaces = append(replaces, "roleName")
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_SOME,
		Replaces:            replaces,
		Stables:             []string{},
		DeleteBeforeReplace: true,
	}, nil
}

func (c *AzureDevopsRoleAssignmentResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputsRoleAssignment := c.ToAzureDevopsRoleAssignmentInput(inputs)
	roleAssignmentId, err := c.setRoleAssignment(inputsRoleAssignment)
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
		Id:         *roleAssignmentId,
		Properties: outputProperties,
	}, nil
}

func (c *AzureDevopsRoleAssignmentResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	var input AzureDevopsRoleAssignmentId
	err := json.Unmarshal([]byte(req.Id), &input)
	if err != nil {
		return nil, err
	}

	return &pbempty.Empty{}, c.removeRoleAssignment(input)
}

func (k *AzureDevopsRoleAssignmentResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (k *AzureDevopsRoleAssignmentResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (k *AzureDevopsRoleAssignmentResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "read is not yet implemented for "+k.Name())
}

func (r *AzureDevopsRoleAssignmentResource) ToAzureDevopsRoleAssignmentInput(inputMap resource.PropertyMap) AzureDevopsRoleAssignmentInput {
	input := AzureDevopsRoleAssignmentInput{}

	if inputMap["resourceId"].HasValue() && inputMap["resourceId"].IsString() {
		input.Id.ResourceId = inputMap["resourceId"].StringValue()
	}

	if inputMap["identityId"].HasValue() && inputMap["identityId"].IsString() {
		input.Id.IdentityId = inputMap["identityId"].StringValue()
	}

	if inputMap["scopeName"].HasValue() && inputMap["scopeName"].IsString() {
		input.Id.ScopeName = ScopeNameInput(inputMap["scopeName"].StringValue())
	}

	if inputMap["userId"].HasValue() && inputMap["userId"].IsString() {
		input.Id.UserId = inputMap["userId"].StringValue()
	}

	if inputMap["roleName"].HasValue() && inputMap["roleName"].IsString() {
		input.RoleName = RoleNameInput(inputMap["roleName"].StringValue())
	}

	return input
}

func (c *AzureDevopsRoleAssignmentResource) setRoleAssignment(input AzureDevopsRoleAssignmentInput) (*string, error) {
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
		"%s/_apis/securityroles/scopes/%s/roleassignments/resources/%s$%s",
		*urlOrg,
		input.Id.ScopeName.GetScopeId(),
		input.Id.ResourceId,
		input.Id.IdentityId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.1-preview.1").
		SetBody([]RoleAssignment{{
			RoleName: string(input.RoleName),
			UserId:   input.Id.UserId,
		}}).
		Put(url)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf(
			"error creating role assignment [%s, %s, %s, %s, %s, %s, %s]': %s",
			*urlOrg,
			input.Id.ScopeName.GetScopeId(),
			input.Id.ResourceId,
			input.Id.IdentityId,
			input.RoleName,
			input.Id.UserId,
			resp.Status(),
			err)
	}

	id := c.createRoleAssignmentId(input.Id)

	return &id, err
}

func (c *AzureDevopsRoleAssignmentResource) createRoleAssignmentId(input AzureDevopsRoleAssignmentId) string {
	data, err := json.Marshal(input)
	if err != nil {
		log.Fatal("error marshalling role assignment input: ", err)
	}

	return string(data)
}

func (c *AzureDevopsRoleAssignmentResource) removeRoleAssignment(roleAssignmentId AzureDevopsRoleAssignmentId) error {
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
		"%s/_apis/securityroles/scopes/%s/roleassignments/resources/%s$%s",
		*urlOrg,
		roleAssignmentId.ScopeName.GetScopeId(),
		roleAssignmentId.ResourceId,
		roleAssignmentId.IdentityId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.1-preview.1").
		SetBody([]string{roleAssignmentId.UserId}).
		Patch(url)

	if err != nil || resp.StatusCode() != 204 {
		return fmt.Errorf(
			"error creating role assignment [%s, %s, %s, %s, %s, %s]': %s",
			*urlOrg,
			roleAssignmentId.ScopeName.GetScopeId(),
			roleAssignmentId.ResourceId,
			roleAssignmentId.IdentityId,
			roleAssignmentId.UserId,
			resp.Status(),
			err)
	}

	return nil
}
