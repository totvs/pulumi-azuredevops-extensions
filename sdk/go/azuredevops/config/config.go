// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The number of attempts.
func GetNumberOfAttempts(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "azuredevops-extensions:numberOfAttempts")
	if err == nil {
		return v
	}
	return getEnvOrDefault(0, parseEnvInt, "NUMBER_OF_ATTEMPTS").(int)
}

// The url of the Azure DevOps instance which should be used.
func GetOrgServiceUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azuredevops-extensions:orgServiceUrl")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AZDO_ORG_SERVICE_URL").(string)
}

// The personal access token which should be used.
func GetPersonalAccessToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azuredevops-extensions:personalAccessToken")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AZDO_PERSONAL_ACCESS_TOKEN").(string)
}
