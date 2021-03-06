# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = ['RoleAssignmentArgs', 'RoleAssignment']

@pulumi.input_type
class RoleAssignmentArgs:
    def __init__(__self__, *,
                 identity_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 role_name: pulumi.Input['RoleName'],
                 scope_name: pulumi.Input['ScopeName'],
                 user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a RoleAssignment resource.
        :param pulumi.Input[str] identity_id: Id of the identity to assign the role to.
        :param pulumi.Input[str] resource_id: Id of the resource on which the role is to be assigned.
        :param pulumi.Input['RoleName'] role_name: The name of the role assigned.
        :param pulumi.Input['ScopeName'] scope_name: The scope name.
        :param pulumi.Input[str] user_id: Unique id of the user given the role assignment.
        """
        pulumi.set(__self__, "identity_id", identity_id)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "scope_name", scope_name)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Input[str]:
        """
        Id of the identity to assign the role to.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        Id of the resource on which the role is to be assigned.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input['RoleName']:
        """
        The name of the role assigned.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input['RoleName']):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="scopeName")
    def scope_name(self) -> pulumi.Input['ScopeName']:
        """
        The scope name.
        """
        return pulumi.get(self, "scope_name")

    @scope_name.setter
    def scope_name(self, value: pulumi.Input['ScopeName']):
        pulumi.set(self, "scope_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        Unique id of the user given the role assignment.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)


class RoleAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input['RoleName']] = None,
                 scope_name: Optional[pulumi.Input['ScopeName']] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RoleAssignment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_id: Id of the identity to assign the role to.
        :param pulumi.Input[str] resource_id: Id of the resource on which the role is to be assigned.
        :param pulumi.Input['RoleName'] role_name: The name of the role assigned.
        :param pulumi.Input['ScopeName'] scope_name: The scope name.
        :param pulumi.Input[str] user_id: Unique id of the user given the role assignment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RoleAssignment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RoleAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input['RoleName']] = None,
                 scope_name: Optional[pulumi.Input['ScopeName']] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleAssignmentArgs.__new__(RoleAssignmentArgs)

            if identity_id is None and not opts.urn:
                raise TypeError("Missing required property 'identity_id'")
            __props__.__dict__["identity_id"] = identity_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            if scope_name is None and not opts.urn:
                raise TypeError("Missing required property 'scope_name'")
            __props__.__dict__["scope_name"] = scope_name
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(RoleAssignment, __self__).__init__(
            'azuredevops-extensions:index:RoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RoleAssignment':
        """
        Get an existing RoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RoleAssignmentArgs.__new__(RoleAssignmentArgs)

        __props__.__dict__["identity_id"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["role_name"] = None
        __props__.__dict__["scope_name"] = None
        __props__.__dict__["user_id"] = None
        return RoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Output[str]:
        """
        Id of the identity to assign the role to.
        """
        return pulumi.get(self, "identity_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        Id of the resource on which the role is to be assigned (ex projectId).
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output['RoleName']:
        """
        The name of the role assigned.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="scopeName")
    def scope_name(self) -> pulumi.Output['ScopeName']:
        """
        The scope name.
        """
        return pulumi.get(self, "scope_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        Unique id of the user given the role assignment.
        """
        return pulumi.get(self, "user_id")

