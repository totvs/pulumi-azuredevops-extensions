// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleName string

const (
	RoleNameReader        = RoleName("Reader")
	RoleNameAdministrator = RoleName("Administrator")
	RoleNameUser          = RoleName("User")
)

func (RoleName) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleName)(nil)).Elem()
}

func (e RoleName) ToRoleNameOutput() RoleNameOutput {
	return pulumi.ToOutput(e).(RoleNameOutput)
}

func (e RoleName) ToRoleNameOutputWithContext(ctx context.Context) RoleNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleNameOutput)
}

func (e RoleName) ToRoleNamePtrOutput() RoleNamePtrOutput {
	return e.ToRoleNamePtrOutputWithContext(context.Background())
}

func (e RoleName) ToRoleNamePtrOutputWithContext(ctx context.Context) RoleNamePtrOutput {
	return RoleName(e).ToRoleNameOutputWithContext(ctx).ToRoleNamePtrOutputWithContext(ctx)
}

func (e RoleName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoleName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleNameOutput struct{ *pulumi.OutputState }

func (RoleNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleName)(nil)).Elem()
}

func (o RoleNameOutput) ToRoleNameOutput() RoleNameOutput {
	return o
}

func (o RoleNameOutput) ToRoleNameOutputWithContext(ctx context.Context) RoleNameOutput {
	return o
}

func (o RoleNameOutput) ToRoleNamePtrOutput() RoleNamePtrOutput {
	return o.ToRoleNamePtrOutputWithContext(context.Background())
}

func (o RoleNameOutput) ToRoleNamePtrOutputWithContext(ctx context.Context) RoleNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleName) *RoleName {
		return &v
	}).(RoleNamePtrOutput)
}

func (o RoleNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoleNamePtrOutput struct{ *pulumi.OutputState }

func (RoleNamePtrOutput) ElementType() reflect.Type {
	return roleNamePtrType
}

func (o RoleNamePtrOutput) ToRoleNamePtrOutput() RoleNamePtrOutput {
	return o
}

func (o RoleNamePtrOutput) ToRoleNamePtrOutputWithContext(ctx context.Context) RoleNamePtrOutput {
	return o
}

func (o RoleNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoleName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o RoleNamePtrOutput) Elem() RoleNameOutput {
	return o.ApplyT(func(v *RoleName) RoleName {
		var ret RoleName
		if v != nil {
			ret = *v
		}
		return ret
	}).(RoleNameOutput)
}

// RoleNameInput is an input type that accepts RoleNameArgs and RoleNameOutput values.
// You can construct a concrete instance of `RoleNameInput` via:
//
//          RoleNameArgs{...}
type RoleNameInput interface {
	pulumi.Input

	ToRoleNameOutput() RoleNameOutput
	ToRoleNameOutputWithContext(context.Context) RoleNameOutput
}

var roleNamePtrType = reflect.TypeOf((**RoleName)(nil)).Elem()

type RoleNamePtrInput interface {
	pulumi.Input

	ToRoleNamePtrOutput() RoleNamePtrOutput
	ToRoleNamePtrOutputWithContext(context.Context) RoleNamePtrOutput
}

type roleNamePtr string

func RoleNamePtr(v string) RoleNamePtrInput {
	return (*roleNamePtr)(&v)
}

func (*roleNamePtr) ElementType() reflect.Type {
	return roleNamePtrType
}

func (in *roleNamePtr) ToRoleNamePtrOutput() RoleNamePtrOutput {
	return pulumi.ToOutput(in).(RoleNamePtrOutput)
}

func (in *roleNamePtr) ToRoleNamePtrOutputWithContext(ctx context.Context) RoleNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoleNamePtrOutput)
}

type ScopeName string

const (
	ScopeNameVariableGroup   = ScopeName("VariableGroup")
	ScopeNameServiceEndpoint = ScopeName("ServiceEndpoint")
)

func (ScopeName) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeName)(nil)).Elem()
}

func (e ScopeName) ToScopeNameOutput() ScopeNameOutput {
	return pulumi.ToOutput(e).(ScopeNameOutput)
}

func (e ScopeName) ToScopeNameOutputWithContext(ctx context.Context) ScopeNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScopeNameOutput)
}

func (e ScopeName) ToScopeNamePtrOutput() ScopeNamePtrOutput {
	return e.ToScopeNamePtrOutputWithContext(context.Background())
}

func (e ScopeName) ToScopeNamePtrOutputWithContext(ctx context.Context) ScopeNamePtrOutput {
	return ScopeName(e).ToScopeNameOutputWithContext(ctx).ToScopeNamePtrOutputWithContext(ctx)
}

func (e ScopeName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScopeName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScopeNameOutput struct{ *pulumi.OutputState }

func (ScopeNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeName)(nil)).Elem()
}

func (o ScopeNameOutput) ToScopeNameOutput() ScopeNameOutput {
	return o
}

func (o ScopeNameOutput) ToScopeNameOutputWithContext(ctx context.Context) ScopeNameOutput {
	return o
}

func (o ScopeNameOutput) ToScopeNamePtrOutput() ScopeNamePtrOutput {
	return o.ToScopeNamePtrOutputWithContext(context.Background())
}

func (o ScopeNameOutput) ToScopeNamePtrOutputWithContext(ctx context.Context) ScopeNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeName) *ScopeName {
		return &v
	}).(ScopeNamePtrOutput)
}

func (o ScopeNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScopeNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScopeNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScopeNamePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamePtrOutput) ElementType() reflect.Type {
	return scopeNamePtrType
}

func (o ScopeNamePtrOutput) ToScopeNamePtrOutput() ScopeNamePtrOutput {
	return o
}

func (o ScopeNamePtrOutput) ToScopeNamePtrOutputWithContext(ctx context.Context) ScopeNamePtrOutput {
	return o
}

func (o ScopeNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScopeName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ScopeNamePtrOutput) Elem() ScopeNameOutput {
	return o.ApplyT(func(v *ScopeName) ScopeName {
		var ret ScopeName
		if v != nil {
			ret = *v
		}
		return ret
	}).(ScopeNameOutput)
}

// ScopeNameInput is an input type that accepts ScopeNameArgs and ScopeNameOutput values.
// You can construct a concrete instance of `ScopeNameInput` via:
//
//          ScopeNameArgs{...}
type ScopeNameInput interface {
	pulumi.Input

	ToScopeNameOutput() ScopeNameOutput
	ToScopeNameOutputWithContext(context.Context) ScopeNameOutput
}

var scopeNamePtrType = reflect.TypeOf((**ScopeName)(nil)).Elem()

type ScopeNamePtrInput interface {
	pulumi.Input

	ToScopeNamePtrOutput() ScopeNamePtrOutput
	ToScopeNamePtrOutputWithContext(context.Context) ScopeNamePtrOutput
}

type scopeNamePtr string

func ScopeNamePtr(v string) ScopeNamePtrInput {
	return (*scopeNamePtr)(&v)
}

func (*scopeNamePtr) ElementType() reflect.Type {
	return scopeNamePtrType
}

func (in *scopeNamePtr) ToScopeNamePtrOutput() ScopeNamePtrOutput {
	return pulumi.ToOutput(in).(ScopeNamePtrOutput)
}

func (in *scopeNamePtr) ToScopeNamePtrOutputWithContext(ctx context.Context) ScopeNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScopeNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleNameOutput{})
	pulumi.RegisterOutputType(RoleNamePtrOutput{})
	pulumi.RegisterOutputType(ScopeNameOutput{})
	pulumi.RegisterOutputType(ScopeNamePtrOutput{})
}
