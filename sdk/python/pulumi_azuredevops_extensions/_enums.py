# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'RoleName',
    'ScopeName',
]


class RoleName(str, Enum):
    READER = "Reader"
    ADMINISTRATOR = "Administrator"
    USER = "User"


class ScopeName(str, Enum):
    VARIABLE_GROUP = "VariableGroup"
