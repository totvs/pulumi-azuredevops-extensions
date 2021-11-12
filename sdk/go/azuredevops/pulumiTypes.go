// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesResource struct {
	// The resource cluster name.
	ClusterName string `pulumi:"clusterName"`
	// The resource name.
	Name string `pulumi:"name"`
	// The resource namemespace.
	Namespace string `pulumi:"namespace"`
	// The service endpoint id.
	ServiceEndpointId string `pulumi:"serviceEndpointId"`
}

// KubernetesResourceInput is an input type that accepts KubernetesResourceArgs and KubernetesResourceOutput values.
// You can construct a concrete instance of `KubernetesResourceInput` via:
//
//          KubernetesResourceArgs{...}
type KubernetesResourceInput interface {
	pulumi.Input

	ToKubernetesResourceOutput() KubernetesResourceOutput
	ToKubernetesResourceOutputWithContext(context.Context) KubernetesResourceOutput
}

type KubernetesResourceArgs struct {
	// The resource cluster name.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The resource name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource namemespace.
	Namespace pulumi.StringInput `pulumi:"namespace"`
	// The service endpoint id.
	ServiceEndpointId pulumi.StringInput `pulumi:"serviceEndpointId"`
}

func (KubernetesResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesResource)(nil)).Elem()
}

func (i KubernetesResourceArgs) ToKubernetesResourceOutput() KubernetesResourceOutput {
	return i.ToKubernetesResourceOutputWithContext(context.Background())
}

func (i KubernetesResourceArgs) ToKubernetesResourceOutputWithContext(ctx context.Context) KubernetesResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesResourceOutput)
}

// KubernetesResourceArrayInput is an input type that accepts KubernetesResourceArray and KubernetesResourceArrayOutput values.
// You can construct a concrete instance of `KubernetesResourceArrayInput` via:
//
//          KubernetesResourceArray{ KubernetesResourceArgs{...} }
type KubernetesResourceArrayInput interface {
	pulumi.Input

	ToKubernetesResourceArrayOutput() KubernetesResourceArrayOutput
	ToKubernetesResourceArrayOutputWithContext(context.Context) KubernetesResourceArrayOutput
}

type KubernetesResourceArray []KubernetesResourceInput

func (KubernetesResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesResource)(nil)).Elem()
}

func (i KubernetesResourceArray) ToKubernetesResourceArrayOutput() KubernetesResourceArrayOutput {
	return i.ToKubernetesResourceArrayOutputWithContext(context.Background())
}

func (i KubernetesResourceArray) ToKubernetesResourceArrayOutputWithContext(ctx context.Context) KubernetesResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesResourceArrayOutput)
}

type KubernetesResourceOutput struct{ *pulumi.OutputState }

func (KubernetesResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesResource)(nil)).Elem()
}

func (o KubernetesResourceOutput) ToKubernetesResourceOutput() KubernetesResourceOutput {
	return o
}

func (o KubernetesResourceOutput) ToKubernetesResourceOutputWithContext(ctx context.Context) KubernetesResourceOutput {
	return o
}

// The resource cluster name.
func (o KubernetesResourceOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResource) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The resource name.
func (o KubernetesResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResource) string { return v.Name }).(pulumi.StringOutput)
}

// The resource namemespace.
func (o KubernetesResourceOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResource) string { return v.Namespace }).(pulumi.StringOutput)
}

// The service endpoint id.
func (o KubernetesResourceOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResource) string { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

type KubernetesResourceArrayOutput struct{ *pulumi.OutputState }

func (KubernetesResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesResource)(nil)).Elem()
}

func (o KubernetesResourceArrayOutput) ToKubernetesResourceArrayOutput() KubernetesResourceArrayOutput {
	return o
}

func (o KubernetesResourceArrayOutput) ToKubernetesResourceArrayOutputWithContext(ctx context.Context) KubernetesResourceArrayOutput {
	return o
}

func (o KubernetesResourceArrayOutput) Index(i pulumi.IntInput) KubernetesResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesResource {
		return vs[0].([]KubernetesResource)[vs[1].(int)]
	}).(KubernetesResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesResourceOutput{})
	pulumi.RegisterOutputType(KubernetesResourceArrayOutput{})
}
