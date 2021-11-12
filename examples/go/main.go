package main

import (
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	aze "github.com/totvs/pulumi-azuredevops-extensions/sdk/v3/go/azuredevops"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		clusteredConventionedName1 := "service-connection-endpoint-kubernetes-pulumi-1"
		clusteredConventionedName2 := "service-connection-endpoint-kubernetes-pulumi-2"
		serviceUrlEndpoint := "https://127.0.0.1:42343"
		projectName := "TOTVSApps"

		project, err := azuredevops.LookupProject(
			ctx,
			&azuredevops.LookupProjectArgs{
				Name: &projectName,
			},
		)

		if err != nil {
			return err
		}

		serviceEndpoint1, err := azuredevops.NewServiceEndpointKubernetes(ctx, clusteredConventionedName1, &azuredevops.ServiceEndpointKubernetesArgs{
			ServiceEndpointName: pulumi.String(clusteredConventionedName1),
			ApiserverUrl:        pulumi.String(serviceUrlEndpoint),
			AuthorizationType:   pulumi.String("Kubeconfig"),
			ProjectId:           pulumi.String(project.Id),
			Kubeconfigs: azuredevops.ServiceEndpointKubernetesKubeconfigArray{
				azuredevops.ServiceEndpointKubernetesKubeconfigArgs{
					KubeConfig: pulumi.String(kubeconfig),
				},
			},
		})

		if err != nil {
			return err
		}

		serviceEndpoint2, err := azuredevops.NewServiceEndpointKubernetes(ctx, clusteredConventionedName2, &azuredevops.ServiceEndpointKubernetesArgs{
			ServiceEndpointName: pulumi.String(clusteredConventionedName2),
			ApiserverUrl:        pulumi.String(serviceUrlEndpoint),
			AuthorizationType:   pulumi.String("Kubeconfig"),
			ProjectId:           pulumi.String(project.Id),
			Kubeconfigs: azuredevops.ServiceEndpointKubernetesKubeconfigArray{
				azuredevops.ServiceEndpointKubernetesKubeconfigArgs{
					KubeConfig: pulumi.String(kubeconfig),
				},
			},
		})

		if err != nil {
			return err
		}

		_, err = aze.NewPipelineEnvironment(ctx, "my-environment", &aze.PipelineEnvironmentArgs{
			Name:      pulumi.String("minha-environment-teste"),
			ProjectId: pulumi.String(projectName),
			KubernetesResources: aze.KubernetesResourceArray{
				aze.KubernetesResourceArgs{
					Name:              pulumi.String("resource-name1"),
					ClusterName:       pulumi.String("cluster-name1"),
					Namespace:         pulumi.String("namespace-name1"),
					ServiceEndpointId: serviceEndpoint1.ID(),
				},
				aze.KubernetesResourceArgs{
					Name:              pulumi.String("resource-name2"),
					ClusterName:       pulumi.String("cluster-name2"),
					Namespace:         pulumi.String("namespace-name2"),
					ServiceEndpointId: serviceEndpoint2.ID(),
				},
			},
		})

		return err
	})
}
