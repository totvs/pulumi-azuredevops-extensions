// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class PipelineEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing PipelineEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PipelineEnvironment {
        return new PipelineEnvironment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops-extensions:index:PipelineEnvironment';

    /**
     * Returns true if the given object is an instance of PipelineEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineEnvironment.__pulumiType;
    }

    /**
     * List of kubernetes resources.
     */
    public readonly kubernetesResources!: pulumi.Output<outputs.KubernetesResource[] | undefined>;
    /**
     * The environment name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a PipelineEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineEnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["kubernetesResources"] = args ? args.kubernetesResources : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
        } else {
            inputs["kubernetesResources"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["projectId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PipelineEnvironment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PipelineEnvironment resource.
 */
export interface PipelineEnvironmentArgs {
    /**
     * List of kubernetes resources.
     */
    kubernetesResources?: pulumi.Input<pulumi.Input<inputs.KubernetesResourceArgs>[]>;
    name: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    projectId: pulumi.Input<string>;
}
