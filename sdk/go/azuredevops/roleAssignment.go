// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleAssignment struct {
	pulumi.CustomResourceState

	// Id of the identity to assign the role to.
	IdentityId pulumi.StringOutput `pulumi:"identityId"`
	// Id of the resource on which the role is to be assigned (ex projectId).
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The name of the role assigned.
	RoleName RoleNameOutput `pulumi:"roleName"`
	// The scope name.
	ScopeName ScopeNameOutput `pulumi:"scopeName"`
	// Unique id of the user given the role assignment.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityId == nil {
		return nil, errors.New("invalid value for required argument 'IdentityId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.ScopeName == nil {
		return nil, errors.New("invalid value for required argument 'ScopeName'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource RoleAssignment
	err := ctx.RegisterResource("azuredevops-extensions:index:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azuredevops-extensions:index:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
}

type RoleAssignmentState struct {
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// Id of the identity to assign the role to.
	IdentityId string `pulumi:"identityId"`
	// Id of the resource on which the role is to be assigned.
	ResourceId string `pulumi:"resourceId"`
	// The name of the role assigned.
	RoleName RoleName `pulumi:"roleName"`
	// The scope name.
	ScopeName ScopeName `pulumi:"scopeName"`
	// Unique id of the user given the role assignment.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// Id of the identity to assign the role to.
	IdentityId pulumi.StringInput
	// Id of the resource on which the role is to be assigned.
	ResourceId pulumi.StringInput
	// The name of the role assigned.
	RoleName RoleNameInput
	// The scope name.
	ScopeName ScopeNameInput
	// Unique id of the user given the role assignment.
	UserId pulumi.StringInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignment)(nil))
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

func (i *RoleAssignment) ToRoleAssignmentPtrOutput() RoleAssignmentPtrOutput {
	return i.ToRoleAssignmentPtrOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentPtrOutputWithContext(ctx context.Context) RoleAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentPtrOutput)
}

type RoleAssignmentPtrInput interface {
	pulumi.Input

	ToRoleAssignmentPtrOutput() RoleAssignmentPtrOutput
	ToRoleAssignmentPtrOutputWithContext(ctx context.Context) RoleAssignmentPtrOutput
}

type roleAssignmentPtrType RoleAssignmentArgs

func (*roleAssignmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil))
}

func (i *roleAssignmentPtrType) ToRoleAssignmentPtrOutput() RoleAssignmentPtrOutput {
	return i.ToRoleAssignmentPtrOutputWithContext(context.Background())
}

func (i *roleAssignmentPtrType) ToRoleAssignmentPtrOutputWithContext(ctx context.Context) RoleAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentPtrOutput)
}

// RoleAssignmentArrayInput is an input type that accepts RoleAssignmentArray and RoleAssignmentArrayOutput values.
// You can construct a concrete instance of `RoleAssignmentArrayInput` via:
//
//          RoleAssignmentArray{ RoleAssignmentArgs{...} }
type RoleAssignmentArrayInput interface {
	pulumi.Input

	ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput
	ToRoleAssignmentArrayOutputWithContext(context.Context) RoleAssignmentArrayOutput
}

type RoleAssignmentArray []RoleAssignmentInput

func (RoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return i.ToRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentArrayOutput)
}

// RoleAssignmentMapInput is an input type that accepts RoleAssignmentMap and RoleAssignmentMapOutput values.
// You can construct a concrete instance of `RoleAssignmentMapInput` via:
//
//          RoleAssignmentMap{ "key": RoleAssignmentArgs{...} }
type RoleAssignmentMapInput interface {
	pulumi.Input

	ToRoleAssignmentMapOutput() RoleAssignmentMapOutput
	ToRoleAssignmentMapOutputWithContext(context.Context) RoleAssignmentMapOutput
}

type RoleAssignmentMap map[string]RoleAssignmentInput

func (RoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return i.ToRoleAssignmentMapOutputWithContext(context.Background())
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentMapOutput)
}

type RoleAssignmentOutput struct {
	*pulumi.OutputState
}

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignment)(nil))
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentPtrOutput() RoleAssignmentPtrOutput {
	return o.ToRoleAssignmentPtrOutputWithContext(context.Background())
}

func (o RoleAssignmentOutput) ToRoleAssignmentPtrOutputWithContext(ctx context.Context) RoleAssignmentPtrOutput {
	return o.ApplyT(func(v RoleAssignment) *RoleAssignment {
		return &v
	}).(RoleAssignmentPtrOutput)
}

type RoleAssignmentPtrOutput struct {
	*pulumi.OutputState
}

func (RoleAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil))
}

func (o RoleAssignmentPtrOutput) ToRoleAssignmentPtrOutput() RoleAssignmentPtrOutput {
	return o
}

func (o RoleAssignmentPtrOutput) ToRoleAssignmentPtrOutputWithContext(ctx context.Context) RoleAssignmentPtrOutput {
	return o
}

type RoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (RoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleAssignment)(nil))
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) Index(i pulumi.IntInput) RoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleAssignment {
		return vs[0].([]RoleAssignment)[vs[1].(int)]
	}).(RoleAssignmentOutput)
}

type RoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (RoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RoleAssignment)(nil))
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) RoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RoleAssignment {
		return vs[0].(map[string]RoleAssignment)[vs[1].(string)]
	}).(RoleAssignmentOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
	pulumi.RegisterOutputType(RoleAssignmentPtrOutput{})
	pulumi.RegisterOutputType(RoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(RoleAssignmentMapOutput{})
}
