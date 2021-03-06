// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevopsExtensions
{
    [AzureDevopsExtensionsResourceType("azuredevops-extensions:index:RoleAssignment")]
    public partial class RoleAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the identity to assign the role to.
        /// </summary>
        [Output("identityId")]
        public Output<string> IdentityId { get; private set; } = null!;

        /// <summary>
        /// Id of the resource on which the role is to be assigned (ex projectId).
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The name of the role assigned.
        /// </summary>
        [Output("roleName")]
        public Output<Pulumi.AzureDevopsExtensions.RoleName> RoleName { get; private set; } = null!;

        /// <summary>
        /// The scope name.
        /// </summary>
        [Output("scopeName")]
        public Output<Pulumi.AzureDevopsExtensions.ScopeName> ScopeName { get; private set; } = null!;

        /// <summary>
        /// Unique id of the user given the role assignment.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a RoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleAssignment(string name, RoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:RoleAssignment", name, args ?? new RoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }
        internal RoleAssignment(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:RoleAssignment", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private RoleAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azuredevops-extensions:index:RoleAssignment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RoleAssignment(name, id, options);
        }
    }

    public sealed class RoleAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the identity to assign the role to.
        /// </summary>
        [Input("identityId", required: true)]
        public Input<string> IdentityId { get; set; } = null!;

        /// <summary>
        /// Id of the resource on which the role is to be assigned.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The name of the role assigned.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<Pulumi.AzureDevopsExtensions.RoleName> RoleName { get; set; } = null!;

        /// <summary>
        /// The scope name.
        /// </summary>
        [Input("scopeName", required: true)]
        public Input<Pulumi.AzureDevopsExtensions.ScopeName> ScopeName { get; set; } = null!;

        /// <summary>
        /// Unique id of the user given the role assignment.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public RoleAssignmentArgs()
        {
        }
    }
}
