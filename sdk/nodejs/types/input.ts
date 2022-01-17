// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

export interface KubernetesResourceArgs {
    /**
     * The resource cluster name.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The resource name.
     */
    name: pulumi.Input<string>;
    /**
     * The resource namemespace.
     */
    namespace: pulumi.Input<string>;
    /**
     * The service endpoint id.
     */
    serviceEndpointId: pulumi.Input<string>;
}
