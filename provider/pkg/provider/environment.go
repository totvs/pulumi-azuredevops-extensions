package provider

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AzureDevopsEnvironmentResource struct {
	config             AzureDevopsConfig
	amountOfTrial      int
	exponentialBackoff time.Duration
}

type AzureDevopsEnvironmentInput struct {
	Name                string
	ProjectId           string
	KubernetesResources []KubernetesResourceInput
}

type KubernetesResourceInput struct {
	Name              string
	ClusterName       string
	Namespace         string
	ServiceEndpointId string
}

func (c *AzureDevopsEnvironmentResource) Name() string {
	return "azuredevops-extensions:index:PipelineEnvironment"
}

func (c *AzureDevopsEnvironmentResource) Configure(config AzureDevopsConfig) {
	c.config = config
}

func (c *AzureDevopsEnvironmentResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
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
	if diffs.Changed("projectId") {
		replaces = append(replaces, "projectId")
	}
	if diffs.Changed("kubernetesResources") {
		replaces = append(replaces, "kubernetesResources")
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_SOME,
		Replaces:            replaces,
		Stables:             []string{},
		DeleteBeforeReplace: true,
	}, nil
}

func (azer *AzureDevopsEnvironmentResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	azer.amountOfTrial = 0
	azer.exponentialBackoff = 1

	numberOfAttempts, err := azer.config.getNumberOfAttempts()
	if err != nil {
		return nil, err
	}

	inputsEnviroment := azer.ToAzureDevopsEnviromentInput(inputs)
	environmentId, err := azer.createEnvironmentPipeline(inputsEnviroment, *numberOfAttempts)
	if err != nil {
		return nil, fmt.Errorf("error creating enviroment pipeline [%s/%s]: %s", inputsEnviroment.ProjectId, inputsEnviroment.Name, err.Error())
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
		Id:         strconv.Itoa(*environmentId),
		Properties: outputProperties,
	}, nil
}

func (c *AzureDevopsEnvironmentResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputsEnviroment := c.ToAzureDevopsEnviromentInput(inputs["__inputs"].ObjectValue())
	environmentId, err := strconv.Atoi(req.Id)
	if err != nil {
		return &pbempty.Empty{}, err
	}

	return &pbempty.Empty{}, c.removeEnvironmentPipeline(environmentId, inputsEnviroment.ProjectId)
}

func (k *AzureDevopsEnvironmentResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (k *AzureDevopsEnvironmentResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	inputsOld, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}
	inputsNew, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	oldEnvironment := k.ToAzureDevopsEnviromentInput(inputsOld["__inputs"].ObjectValue())
	newEnvironment := k.ToAzureDevopsEnviromentInput(inputsNew)
	environmentId, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, fmt.Errorf("error parsing enviroment to int [%s/%s]: %s", newEnvironment.ProjectId, req.Id, err.Error())
	}

	err = k.updateEnvironmentPipeline(
		environmentId,
		oldEnvironment,
		newEnvironment,
	)
	if err != nil {
		return nil, err
	}

	outputStore := resource.PropertyMap{}
	outputStore["__inputs"] = resource.NewObjectProperty(inputsNew)

	outputProperties, err := plugin.MarshalProperties(
		outputStore,
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

func (k *AzureDevopsEnvironmentResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "read is not yet implemented for "+k.Name())
}

func (r *AzureDevopsEnvironmentResource) ToAzureDevopsEnviromentInput(inputMap resource.PropertyMap) AzureDevopsEnvironmentInput {
	input := AzureDevopsEnvironmentInput{}

	if inputMap["name"].HasValue() && inputMap["name"].IsString() {
		input.Name = inputMap["name"].StringValue()
	}

	if inputMap["projectId"].HasValue() && inputMap["projectId"].IsString() {
		input.ProjectId = inputMap["projectId"].StringValue()
	}

	if inputMap["kubernetesResources"].HasValue() && inputMap["kubernetesResources"].IsArray() {
		for _, m := range inputMap["kubernetesResources"].ArrayValue() {
			if m.HasValue() && m.IsObject() {
				var kr KubernetesResourceInput
				if m.ObjectValue()["name"].HasValue() {
					kr.Name = m.ObjectValue()["name"].StringValue()
				}
				if m.ObjectValue()["namespace"].HasValue() {
					kr.Namespace = m.ObjectValue()["namespace"].StringValue()
				}
				if m.ObjectValue()["clusterName"].HasValue() {
					kr.ClusterName = m.ObjectValue()["clusterName"].StringValue()
				}
				if m.ObjectValue()["serviceEndpointId"].HasValue() {
					kr.ServiceEndpointId = m.ObjectValue()["serviceEndpointId"].StringValue()
				}

				input.KubernetesResources = append(input.KubernetesResources, kr)
			}
		}
	}

	return input
}

/// https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-6.0
func (c *AzureDevopsEnvironmentResource) createEnvironmentPipeline(input AzureDevopsEnvironmentInput, numberOfAttempts int) (*int, error) {

	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return nil, err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments", *urlOrg, input.ProjectId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.1-preview.1").
		SetBody(map[string]interface{}{"name": input.Name}).
		Post(url)

	if err != nil || resp.StatusCode() != 200 {

		if c.amountOfTrial < numberOfAttempts {
			c.amountOfTrial++

			c.exponentialBackoff *= 2
			fmt.Printf("try #%s, Next check will be after %s seconds", c.amountOfTrial, c.exponentialBackoff.Seconds())
			time.Sleep(c.exponentialBackoff)

			return c.createEnvironmentPipeline(input, numberOfAttempts)
		}

		return nil, fmt.Errorf("error creating enviroment pipeline [%s/%s/%s]': %s", input.ProjectId, input.Name, resp.Status(), err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(resp.Body()), &result)
	id := int(result["id"].(float64))

	for _, resource := range input.KubernetesResources {
		_, err = c.createResourceEnvironmentPipeline(
			resource.Name,
			input.ProjectId,
			id,
			resource.ClusterName,
			resource.Namespace,
			resource.ServiceEndpointId,
		)

		if err != nil {
			c.removeEnvironmentPipeline(id, input.ProjectId)
			return nil, fmt.Errorf("error creating environment pipeline [%s/%d]: %s", input.ProjectId, id, err.Error())
		}
	}

	return &id, err
}

/// https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/kubernetes?view=azure-devops-rest-6.0
func (c *AzureDevopsEnvironmentResource) createResourceEnvironmentPipeline(
	name string,
	projectId string,
	environmentId int,
	clusterName string,
	namespace string,
	serviceEndpointId string) (*int, error) {

	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return nil, err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d/providers/kubernetes", *urlOrg, projectId, environmentId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.1-preview.1").
		SetBody(map[string]interface{}{
			"name":              name,
			"clusterName":       clusterName,
			"namespace":         namespace,
			"serviceEndpointId": serviceEndpointId,
		}).
		Post(url)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error creating resource environment pipeline [%s/%s/%s/%s]: %s", projectId, serviceEndpointId, name, resp.Status(), err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(resp.Body()), &result)
	id := int(result["id"].(float64))

	return &id, err
}

func (c *AzureDevopsEnvironmentResource) removeEnvironmentPipeline(environmentId int, projectId string) error {
	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return err
	}

	client := resty.New()
	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d", *urlOrg, projectId, environmentId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.0-preview.1").
		Delete(url)

	if err != nil || resp.StatusCode() != 204 {
		return fmt.Errorf("error removing enviroment pipeline [%s/%d/%s]: %s", projectId, environmentId, resp.Status(), err)
	}

	return nil
}

func (c *AzureDevopsEnvironmentResource) updateEnvironmentPipeline(
	environmentId int,
	old AzureDevopsEnvironmentInput,
	new AzureDevopsEnvironmentInput) (err error) {

	urlOrg, err := c.config.getOrgServiceUrl()
	if err != nil {
		return err
	}

	pat, err := c.config.getPersonalAccessToken()
	if err != nil {
		return err
	}

	client := resty.New()
	url := fmt.Sprintf("%s/%s/_apis/distributedtask/environments/%d?api-version=6.0-preview.1", *urlOrg, new.ProjectId, environmentId)
	resp, err := client.R().
		SetBasicAuth("pat", *pat).
		SetQueryString("api-version=6.0-preview.1").
		SetBody(map[string]interface{}{
			"name": new.Name,
		}).
		Patch(url)

	if err != nil || resp.StatusCode() != 200 {
		return fmt.Errorf("error updating enviroment [%s/%d/%s]: %s", new.ProjectId, environmentId, resp.Status(), err)
	}

	return nil
}
