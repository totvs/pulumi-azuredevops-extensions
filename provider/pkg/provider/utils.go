package provider

import (
	"fmt"
	"os"
	"strconv"

	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

var AzureDevopsResources = []AzureDevopsResource{
	&AzureDevopsEnvironmentResource{},
	&AzureDevopsRoleAssignmentResource{},
}

type AzureDevopsResource interface {
	Name() string
	Configure(config AzureDevopsConfig)
	Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error)
	Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error)
	Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error)
	Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error)
	Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error)
	Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error)
}

type ResourceBase interface {
	GetUrn() string
}

type AzureDevopsConfig struct {
	Config map[string]string
}

func (sc *AzureDevopsConfig) getConfig(configName, envName string) string {
	if val, ok := sc.Config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

func (sc *AzureDevopsConfig) getOrgServiceUrl() (*string, error) {
	token := sc.getConfig("azuredevops-extensions:config:orgServiceUrl", "AZDO_ORG_SERVICE_URL")

	if len(token) == 0 {
		return nil, fmt.Errorf("no org for azure devops service url found")
	}

	return &token, nil
}

func (sc *AzureDevopsConfig) getPersonalAccessToken() (*string, error) {
	token := sc.getConfig("azuredevops-extensions:config:personalAccessToken", "AZDO_PERSONAL_ACCESS_TOKEN")

	if len(token) == 0 {
		return nil, fmt.Errorf("no personal access token for azure devops found")
	}

	return &token, nil
}

func getAzureDevopsResource(name string) AzureDevopsResource {
	for _, r := range AzureDevopsResources {
		if r.Name() == name {
			return r
		}
	}

	return &AzureDevopsUnknownResource{}
}

func getResourceNameFromRequest(req ResourceBase) string {
	urn := resource.URN(req.GetUrn())
	return urn.Type().String()
}

func (sc *AzureDevopsConfig) getNumberOfAttempts() (*int, error) {
	token := sc.getConfig("azuredevops-extensions:config:numberOfAttempts", "NUMBER_OF_ATTEMPTS")

	numberOfAttempts, err := strconv.Atoi("0")

	if err != nil {
		return nil, err
	}

	if len(token) == 0 {
		return &numberOfAttempts, nil
	}

	numberOfAttempts, err = strconv.Atoi(token)

	if err != nil {
		return nil, err
	}

	return &numberOfAttempts, nil
}
