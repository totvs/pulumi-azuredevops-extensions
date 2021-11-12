# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'KubernetesResource',
]

@pulumi.output_type
class KubernetesResource(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterName":
            suggest = "cluster_name"
        elif key == "serviceEndpointId":
            suggest = "service_endpoint_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KubernetesResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KubernetesResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KubernetesResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_name: str,
                 name: str,
                 namespace: str,
                 service_endpoint_id: str):
        """
        :param str cluster_name: The resource cluster name.
        :param str name: The resource name.
        :param str namespace: The resource namemespace.
        :param str service_endpoint_id: The service endpoint id.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        The resource cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        The resource namemespace.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> str:
        """
        The service endpoint id.
        """
        return pulumi.get(self, "service_endpoint_id")


