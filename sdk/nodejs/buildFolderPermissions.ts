// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for an Folder (Component)
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Permission levels
 *
 * Permission for Areas within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const project-readers = project.id.apply(id => azuredevops.getGroup({
 *     projectId: id,
 *     name: "Readers",
 * }));
 * const root_permissions = new azuredevops.AreaPermissions("root-permissions", {
 *     projectId: project.id,
 *     principal: project_readers.apply(project_readers => project_readers.id),
 *     path: "/",
 *     permissions: {
 *         CREATE_CHILDREN: "Deny",
 *         GENERIC_READ: "Allow",
 *         DELETE: "Deny",
 *         WORK_ITEM_READ: "Allow",
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class BuildFolderPermissions extends pulumi.CustomResource {
    /**
     * Get an existing BuildFolderPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildFolderPermissionsState, opts?: pulumi.CustomResourceOptions): BuildFolderPermissions {
        return new BuildFolderPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops-extensions:index:BuildFolderPermissions';

    /**
     * Returns true if the given object is an instance of BuildFolderPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BuildFolderPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BuildFolderPermissions.__pulumiType;
    }

    /**
     * The name of the folder to assign the permissions.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a BuildFolderPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildFolderPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BuildFolderPermissionsArgs | BuildFolderPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildFolderPermissionsState | undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["principal"] = state ? state.principal : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as BuildFolderPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["path"] = args ? args.path : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["principal"] = args ? args.principal : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["replace"] = args ? args.replace : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BuildFolderPermissions.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FolderPermissions resources.
 */
export interface BuildFolderPermissionsState {
    /**
     * The name of the branch to assign the permissions.
     */
    path?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a BuildFolderPermissions resource.
 */
export interface BuildFolderPermissionsArgs {
    /**
     * The name of the branch to assign the permissions.
     */
    path?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    replace?: pulumi.Input<boolean>;
}
