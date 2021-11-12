module test

go 1.14

require (
	github.com/pulumi/pulumi-azuredevops/sdk/v2 v2.2.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
	github.com/totvs/pulumi-azuredevops-extensions/sdk/v3 v3.0.0
)

replace github.com/totvs/pulumi-azuredevops-extensions/sdk/v3 => /git/go/src/github.com/totvs/pulumi-azuredevops-extensions/sdk
