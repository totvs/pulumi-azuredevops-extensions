package provider

import (
	"fmt"

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

type AzureDevopsRoleAssignmentInput struct {
	ResourceId string
	IdentityId string
	ScopeName  ScopeNameInput
	UserId     string
	RoleName   RoleNameInput
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
	return nil, nil
	// olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: false, SkipNulls: true})
	// if err != nil {
	// 	return nil, err
	// }

	// news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: false})
	// if err != nil {
	// 	return nil, err
	// }

	// diffs := olds["__inputs"].ObjectValue().Diff(news)
	// if diffs == nil {
	// 	return &pulumirpc.DiffResponse{
	// 		Changes:             pulumirpc.DiffResponse_DIFF_NONE,
	// 		Replaces:            []string{},
	// 		Stables:             []string{},
	// 		DeleteBeforeReplace: false,
	// 	}, nil
	// }

	// var replaces []string
	// if diffs.Changed("projectId") {
	// 	replaces = append(replaces, "projectId")
	// }
	// if diffs.Changed("kubernetesResources") {
	// 	replaces = append(replaces, "kubernetesResources")
	// }

	// return &pulumirpc.DiffResponse{
	// 	Changes:             pulumirpc.DiffResponse_DIFF_SOME,
	// 	Replaces:            replaces,
	// 	Stables:             []string{},
	// 	DeleteBeforeReplace: true,
	// }, nil
}

func (c *AzureDevopsRoleAssignmentResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	numberOfAttempts, err := c.config.getNumberOfAttempts()
	if err != nil {
		return nil, err
	}

	inputsRoleAssignment := c.ToAzureDevopsRoleAssignmentInput(inputs)
	roleAssignmentId, err := c.createRoleAssignment(inputsRoleAssignment, *numberOfAttempts)
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
	return nil, nil
	// inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	// if err != nil {
	// 	return nil, err
	// }

	// inputsEnviroment := c.ToAzureDevopsRoleAssignmentInput(inputs["__inputs"].ObjectValue())
	// environmentId, err := strconv.Atoi(req.Id)
	// if err != nil {
	// 	return &pbempty.Empty{}, err
	// }

	// return &pbempty.Empty{}, c.removeEnvironmentPipeline(environmentId, inputsEnviroment.ProjectId)
}

func (k *AzureDevopsRoleAssignmentResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (k *AzureDevopsRoleAssignmentResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, nil
	// inputsOld, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	// if err != nil {
	// 	return nil, err
	// }
	// inputsNew, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	// if err != nil {
	// 	return nil, err
	// }

	// oldEnvironment := k.ToAzureDevopsRoleAssignmentInput(inputsOld["__inputs"].ObjectValue())
	// newEnvironment := k.ToAzureDevopsRoleAssignmentInput(inputsNew)
	// environmentId, err := strconv.Atoi(req.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("error parsing enviroment to int [%s/%s]: %s", newEnvironment.ProjectId, req.Id, err.Error())
	// }

	// err = k.updateEnvironmentPipeline(
	// 	environmentId,
	// 	oldEnvironment,
	// 	newEnvironment,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// outputStore := resource.PropertyMap{}
	// outputStore["__inputs"] = resource.NewObjectProperty(inputsNew)

	// outputProperties, err := plugin.MarshalProperties(
	// 	outputStore,
	// 	plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// return &pulumirpc.UpdateResponse{
	// 	Properties: outputProperties,
	// }, nil
}

func (k *AzureDevopsRoleAssignmentResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "read is not yet implemented for "+k.Name())
}

func (r *AzureDevopsRoleAssignmentResource) ToAzureDevopsRoleAssignmentInput(inputMap resource.PropertyMap) AzureDevopsRoleAssignmentInput {
	input := AzureDevopsRoleAssignmentInput{}

	if inputMap["resourceId"].HasValue() && inputMap["resourceId"].IsString() {
		input.ResourceId = inputMap["resourceId"].StringValue()
	}

	if inputMap["identityId"].HasValue() && inputMap["identityId"].IsString() {
		input.IdentityId = inputMap["identityId"].StringValue()
	}

	if inputMap["scopeName"].HasValue() && inputMap["scopeName"].IsString() {
		input.ScopeName = ScopeNameInput(inputMap["scopeName"].StringValue())
	}

	if inputMap["userId"].HasValue() && inputMap["userId"].IsString() {
		input.UserId = inputMap["userId"].StringValue()
	}

	if inputMap["roleName"].HasValue() && inputMap["roleName"].IsString() {
		input.RoleName = RoleNameInput(inputMap["roleName"].StringValue())
	}

	return input
}

func (c *AzureDevopsRoleAssignmentResource) createRoleAssignment(input AzureDevopsRoleAssignmentInput, numberOfAttempts int) (*string, error) {

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
		"%s/_apis/securityroles/scopes/%s/roleassignments/resources/%s/%s",
		*urlOrg,
		input.ScopeName.GetScopeId(),
		input.ResourceId,
		input.IdentityId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.1-preview.1").
		SetBody(map[string]interface{}{
			"roleName": input.RoleName,
			"userId":   input.UserId,
		}).
		Put(url)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf(
			"error creating role assignment [%s, %s, %s, %s, %s/%s/%s]': %s",
			*urlOrg,
			input.ScopeName.GetScopeId(),
			input.ResourceId,
			input.IdentityId,
			input.RoleName,
			input.UserId,
			resp.Status(),
			err)
	}

	id := c.createRoleAssignmentId(input.RoleName, input.UserId)

	return &id, err
}

func (c *AzureDevopsRoleAssignmentResource) createRoleAssignmentId(roleName RoleNameInput, userId string) string {
	return fmt.Sprintf("%s/%s", roleName, userId)
}

// func (c *AzureDevopsRoleAssignmentsResource) createResourceEnvironmentPipeline(
// 	name string,
// 	projectId string,
// 	environmentId int,
// 	clusterName string,
// 	namespace string,
// 	serviceEndpointId string,
// 	numberOfAttempts int) (*int, error) {

// 	urlOrg, err := c.config.getOrgServiceUrl()
// 	if err != nil {
// 		return nil, err
// 	}

// 	pat, err := c.config.getPersonalAccessToken()
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := resty.New()
// 	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d/providers/kubernetes", *urlOrg, projectId, environmentId)
// 	resp, err := client.R().
// 		SetBasicAuth("pat", *pat).
// 		SetQueryString("api-version=6.1-preview.1").
// 		SetBody(map[string]interface{}{
// 			"name":              name,
// 			"clusterName":       clusterName,
// 			"namespace":         namespace,
// 			"serviceEndpointId": serviceEndpointId,
// 		}).
// 		Post(url)

// 	if err != nil || resp.StatusCode() > 399 {
// 		if c.amountOfTrial < numberOfAttempts {
// 			c.amountOfTrial++

// 			c.exponentialBackoff *= 2
// 			fmt.Printf("try #%d, next attempt on %f seconds\n", c.amountOfTrial, c.exponentialBackoff.Seconds())
// 			time.Sleep(c.exponentialBackoff)

// 			return c.createResourceEnvironmentPipeline(
// 				name,
// 				projectId,
// 				environmentId,
// 				clusterName,
// 				namespace,
// 				serviceEndpointId,
// 				numberOfAttempts,
// 			)
// 		}

// 		return nil, fmt.Errorf("error creating resource environment pipeline [%s/%s/%s/%s]: %s", projectId, serviceEndpointId, name, resp.Status(), err)
// 	}

// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(resp.Body()), &result)
// 	id := int(result["id"].(float64))

// 	return &id, err
// }

// func (c *AzureDevopsRoleAssignmentsResource) removeEnvironmentPipeline(environmentId int, projectId string) error {
// 	urlOrg, err := c.config.getOrgServiceUrl()
// 	if err != nil {
// 		return err
// 	}

// 	pat, err := c.config.getPersonalAccessToken()
// 	if err != nil {
// 		return err
// 	}

// 	client := resty.New()
// 	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d", *urlOrg, projectId, environmentId)
// 	resp, err := client.R().
// 		SetBasicAuth("pat", *pat).
// 		SetQueryString("api-version=6.0-preview.1").
// 		Delete(url)

// 	if err != nil || resp.StatusCode() != 204 {
// 		return fmt.Errorf("error removing enviroment pipeline [%s/%d/%s]: %s", projectId, environmentId, resp.Status(), err)
// 	}

// 	return nil
// }

// func (c *AzureDevopsRoleAssignmentsResource) updateEnvironmentPipeline(
// 	environmentId int,
// 	old AzureDevopsEnvironmentInput,
// 	new AzureDevopsEnvironmentInput) (err error) {

// 	urlOrg, err := c.config.getOrgServiceUrl()
// 	if err != nil {
// 		return err
// 	}

// 	pat, err := c.config.getPersonalAccessToken()
// 	if err != nil {
// 		return err
// 	}

// 	client := resty.New()
// 	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d?api-version=6.0-preview.1", *urlOrg, new.ProjectId, environmentId)
// 	resp, err := client.R().
// 		SetBasicAuth("pat", *pat).
// 		SetQueryString("api-version=6.0-preview.1").
// 		SetBody(map[string]interface{}{
// 			"name": new.Name,
// 		}).
// 		Patch(url)

// 	if err != nil || resp.StatusCode() != 200 {
// 		return fmt.Errorf("error updating enviroment [%s/%d/%s]: %s", new.ProjectId, environmentId, resp.Status(), err)
// 	}

// 	return nil
// }
