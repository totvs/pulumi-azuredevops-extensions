// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevopsExtensions.Inputs
{

    public sealed class KubernetesResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource cluster name.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The resource namemespace.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The service endpoint id.
        /// </summary>
        [Input("serviceEndpointId", required: true)]
        public Input<string> ServiceEndpointId { get; set; } = null!;

        public KubernetesResourceArgs()
        {
        }
    }
}
