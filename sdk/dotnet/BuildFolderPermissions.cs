// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevopsExtensions
{
    /// <summary>
    /// Manages permissions for an Folder (Component)
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Permission levels
    /// 
    /// Permission for Areas within Azure DevOps can be applied on two different levels.
    /// Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
    /// 
    /// ## Example Usage
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
    ///         {
    ///             WorkItemTemplate = "Agile",
    ///             VersionControl = "Git",
    ///             Visibility = "private",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var project_readers = project.Id.Apply(id =&gt; AzureDevOps.GetGroup.InvokeAsync(new AzureDevOps.GetGroupArgs
    ///         {
    ///             ProjectId = id,
    ///             Name = "Readers",
    ///         }));
    ///         var root_permissions = new AzureDevOps.AreaPermissions("root-permissions", new AzureDevOps.AreaPermissionsArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Principal = project_readers.Apply(project_readers =&gt; project_readers.Id),
    ///             Path = "/",
    ///             Permissions = 
    ///             {
    ///                 { "CREATE_CHILDREN", "Deny" },
    ///                 { "GENERIC_READ", "Allow" },
    ///                 { "DELETE", "Deny" },
    ///                 { "WORK_ITEM_READ", "Allow" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
    /// 
    /// ## Import
    /// 
    /// The resource does not support import.
    /// </summary>
    [AzureDevopsExtensionsResourceType("azuredevops-extensions:index:BuildFolderPermissions")]
    public partial class BuildFolderPermissions : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the folder to assign the permissions.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableDictionary<string, string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a BuildFolderPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BuildFolderPermissions(string name, BuildFolderPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:BuildFolderPermissions", name, args ?? new BuildFolderPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }
        internal BuildFolderPermissions(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:BuildFolderPermissions", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private BuildFolderPermissions(string name, Input<string> id, BuildFolderPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:BuildFolderPermissions", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BuildFolderPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BuildFolderPermissions Get(string name, Input<string> id, BuildFolderPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new BuildFolderPermissions(name, id, state, options);
        }
    }

    public sealed class BuildFolderPermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public BuildFolderPermissionsArgs()
        {
        }
    }

    public sealed class BuildFolderPermissionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public BuildFolderPermissionsState()
        {
        }
    }
}
