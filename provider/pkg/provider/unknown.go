package provider

import (
	"fmt"

	pbempty "github.com/golang/protobuf/ptypes/empty"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

type AzureDevopsUnknownResource struct{}

func (c AzureDevopsUnknownResource) Name() string {
	return "azuredevops-extensions:index::Unknown"
}

func (u *AzureDevopsUnknownResource) Configure(config AzureDevopsConfig) {
}

func (c *AzureDevopsUnknownResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (c *AzureDevopsUnknownResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (c *AzureDevopsUnknownResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (k *AzureDevopsUnknownResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (k *AzureDevopsUnknownResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (k *AzureDevopsUnknownResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func createUnknownResourceErrorFromRequest(req ResourceBase) error {
	rn := getResourceNameFromRequest(req)
	return fmt.Errorf("unknown resource type '%s'", rn)
}
