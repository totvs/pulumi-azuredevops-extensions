// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PipelineEnvironment struct {
	pulumi.CustomResourceState

	// List of kubernetes resources.
	KubernetesResources KubernetesResourceArrayOutput `pulumi:"kubernetesResources"`
	// The environment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewPipelineEnvironment registers a new resource with the given unique name, arguments, and options.
func NewPipelineEnvironment(ctx *pulumi.Context,
	name string, args *PipelineEnvironmentArgs, opts ...pulumi.ResourceOption) (*PipelineEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource PipelineEnvironment
	err := ctx.RegisterResource("azuredevops-extensions:index:PipelineEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineEnvironment gets an existing PipelineEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineEnvironmentState, opts ...pulumi.ResourceOption) (*PipelineEnvironment, error) {
	var resource PipelineEnvironment
	err := ctx.ReadResource("azuredevops-extensions:index:PipelineEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineEnvironment resources.
type pipelineEnvironmentState struct {
}

type PipelineEnvironmentState struct {
}

func (PipelineEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineEnvironmentState)(nil)).Elem()
}

type pipelineEnvironmentArgs struct {
	// List of kubernetes resources.
	KubernetesResources []KubernetesResource `pulumi:"kubernetesResources"`
	Name                string               `pulumi:"name"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a PipelineEnvironment resource.
type PipelineEnvironmentArgs struct {
	// List of kubernetes resources.
	KubernetesResources KubernetesResourceArrayInput
	Name                pulumi.StringInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
}

func (PipelineEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineEnvironmentArgs)(nil)).Elem()
}

type PipelineEnvironmentInput interface {
	pulumi.Input

	ToPipelineEnvironmentOutput() PipelineEnvironmentOutput
	ToPipelineEnvironmentOutputWithContext(ctx context.Context) PipelineEnvironmentOutput
}

func (*PipelineEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineEnvironment)(nil))
}

func (i *PipelineEnvironment) ToPipelineEnvironmentOutput() PipelineEnvironmentOutput {
	return i.ToPipelineEnvironmentOutputWithContext(context.Background())
}

func (i *PipelineEnvironment) ToPipelineEnvironmentOutputWithContext(ctx context.Context) PipelineEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineEnvironmentOutput)
}

func (i *PipelineEnvironment) ToPipelineEnvironmentPtrOutput() PipelineEnvironmentPtrOutput {
	return i.ToPipelineEnvironmentPtrOutputWithContext(context.Background())
}

func (i *PipelineEnvironment) ToPipelineEnvironmentPtrOutputWithContext(ctx context.Context) PipelineEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineEnvironmentPtrOutput)
}

type PipelineEnvironmentPtrInput interface {
	pulumi.Input

	ToPipelineEnvironmentPtrOutput() PipelineEnvironmentPtrOutput
	ToPipelineEnvironmentPtrOutputWithContext(ctx context.Context) PipelineEnvironmentPtrOutput
}

type pipelineEnvironmentPtrType PipelineEnvironmentArgs

func (*pipelineEnvironmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineEnvironment)(nil))
}

func (i *pipelineEnvironmentPtrType) ToPipelineEnvironmentPtrOutput() PipelineEnvironmentPtrOutput {
	return i.ToPipelineEnvironmentPtrOutputWithContext(context.Background())
}

func (i *pipelineEnvironmentPtrType) ToPipelineEnvironmentPtrOutputWithContext(ctx context.Context) PipelineEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineEnvironmentPtrOutput)
}

// PipelineEnvironmentArrayInput is an input type that accepts PipelineEnvironmentArray and PipelineEnvironmentArrayOutput values.
// You can construct a concrete instance of `PipelineEnvironmentArrayInput` via:
//
//          PipelineEnvironmentArray{ PipelineEnvironmentArgs{...} }
type PipelineEnvironmentArrayInput interface {
	pulumi.Input

	ToPipelineEnvironmentArrayOutput() PipelineEnvironmentArrayOutput
	ToPipelineEnvironmentArrayOutputWithContext(context.Context) PipelineEnvironmentArrayOutput
}

type PipelineEnvironmentArray []PipelineEnvironmentInput

func (PipelineEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineEnvironment)(nil)).Elem()
}

func (i PipelineEnvironmentArray) ToPipelineEnvironmentArrayOutput() PipelineEnvironmentArrayOutput {
	return i.ToPipelineEnvironmentArrayOutputWithContext(context.Background())
}

func (i PipelineEnvironmentArray) ToPipelineEnvironmentArrayOutputWithContext(ctx context.Context) PipelineEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineEnvironmentArrayOutput)
}

// PipelineEnvironmentMapInput is an input type that accepts PipelineEnvironmentMap and PipelineEnvironmentMapOutput values.
// You can construct a concrete instance of `PipelineEnvironmentMapInput` via:
//
//          PipelineEnvironmentMap{ "key": PipelineEnvironmentArgs{...} }
type PipelineEnvironmentMapInput interface {
	pulumi.Input

	ToPipelineEnvironmentMapOutput() PipelineEnvironmentMapOutput
	ToPipelineEnvironmentMapOutputWithContext(context.Context) PipelineEnvironmentMapOutput
}

type PipelineEnvironmentMap map[string]PipelineEnvironmentInput

func (PipelineEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineEnvironment)(nil)).Elem()
}

func (i PipelineEnvironmentMap) ToPipelineEnvironmentMapOutput() PipelineEnvironmentMapOutput {
	return i.ToPipelineEnvironmentMapOutputWithContext(context.Background())
}

func (i PipelineEnvironmentMap) ToPipelineEnvironmentMapOutputWithContext(ctx context.Context) PipelineEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineEnvironmentMapOutput)
}

type PipelineEnvironmentOutput struct {
	*pulumi.OutputState
}

func (PipelineEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineEnvironment)(nil))
}

func (o PipelineEnvironmentOutput) ToPipelineEnvironmentOutput() PipelineEnvironmentOutput {
	return o
}

func (o PipelineEnvironmentOutput) ToPipelineEnvironmentOutputWithContext(ctx context.Context) PipelineEnvironmentOutput {
	return o
}

func (o PipelineEnvironmentOutput) ToPipelineEnvironmentPtrOutput() PipelineEnvironmentPtrOutput {
	return o.ToPipelineEnvironmentPtrOutputWithContext(context.Background())
}

func (o PipelineEnvironmentOutput) ToPipelineEnvironmentPtrOutputWithContext(ctx context.Context) PipelineEnvironmentPtrOutput {
	return o.ApplyT(func(v PipelineEnvironment) *PipelineEnvironment {
		return &v
	}).(PipelineEnvironmentPtrOutput)
}

type PipelineEnvironmentPtrOutput struct {
	*pulumi.OutputState
}

func (PipelineEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineEnvironment)(nil))
}

func (o PipelineEnvironmentPtrOutput) ToPipelineEnvironmentPtrOutput() PipelineEnvironmentPtrOutput {
	return o
}

func (o PipelineEnvironmentPtrOutput) ToPipelineEnvironmentPtrOutputWithContext(ctx context.Context) PipelineEnvironmentPtrOutput {
	return o
}

type PipelineEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (PipelineEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PipelineEnvironment)(nil))
}

func (o PipelineEnvironmentArrayOutput) ToPipelineEnvironmentArrayOutput() PipelineEnvironmentArrayOutput {
	return o
}

func (o PipelineEnvironmentArrayOutput) ToPipelineEnvironmentArrayOutputWithContext(ctx context.Context) PipelineEnvironmentArrayOutput {
	return o
}

func (o PipelineEnvironmentArrayOutput) Index(i pulumi.IntInput) PipelineEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PipelineEnvironment {
		return vs[0].([]PipelineEnvironment)[vs[1].(int)]
	}).(PipelineEnvironmentOutput)
}

type PipelineEnvironmentMapOutput struct{ *pulumi.OutputState }

func (PipelineEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PipelineEnvironment)(nil))
}

func (o PipelineEnvironmentMapOutput) ToPipelineEnvironmentMapOutput() PipelineEnvironmentMapOutput {
	return o
}

func (o PipelineEnvironmentMapOutput) ToPipelineEnvironmentMapOutputWithContext(ctx context.Context) PipelineEnvironmentMapOutput {
	return o
}

func (o PipelineEnvironmentMapOutput) MapIndex(k pulumi.StringInput) PipelineEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PipelineEnvironment {
		return vs[0].(map[string]PipelineEnvironment)[vs[1].(string)]
	}).(PipelineEnvironmentOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineEnvironmentOutput{})
	pulumi.RegisterOutputType(PipelineEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(PipelineEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(PipelineEnvironmentMapOutput{})
}
