# Intro
This provider was created to complement the original [AzureDevOps Provider](https://www.pulumi.com/registry/packages/azuredevops). The original provider doesn't have resource environments to support pipelines as yaml code. In the future, maybe, we'll have this feature and this provider will be deprecated.

## Permissions

You need to create a new personal access token (PAT) for an Azure DevOps account. See the list of needed permissions below:

```yaml
Service Connections:
- Read
Environment:
- Read & manage
```

## Environment

You can set Azure DevOps token and organization service url with:

```sh
export AZDO_PERSONAL_ACCESS_TOKEN=ahdflakjsdhdfkajsdhdflkajshdflakjs34t9813h4v134ht3i4
export AZDO_ORG_SERVICE_URL=https://dev.azure.com/typeorganizationhere
export NUMBER_OF_ATTEMPTS=3
```
or
```sh
pulumi config set --secret azuredevops-extensions:config:orgServiceUrl https://dev.azure.com/typeorganizationhere
pulumi config set --secret azuredevops-extensions:config:personalAccessToken ahdflakjsdhdfkajsdhdflkajshdflakjs34t9813h4v134ht3i4
pulumi config set --secret azuredevops-extensions:config:numberOfAttempts 3
```

# Installation

You can install the plugin with these commands:

```bash
wget https://github.com/totvs/pulumi-azuredevops-extensions/releases/download/v0.0.2-alpha.1636747691%2Bbb5b83c8/pulumi-resource-azuredevops-extensions-v0.0.2-alpha.1636747691+bb5b83c8-linux-amd64.tar.gz

pulumi plugin install resource azuredevops-extensions 0.0.2 -f pulumi-resource-azuredevops-extensions-v0.0.2-alpha.1636747691+bb5b83c8-linux-amd64.tar.gz
```

# Import

## Build pipeline folder import
To import build pipeline folder, use the following format ```<azuredevops-project-id>\/<folder-name>```. Example:


```sh
pulumi import 'azuredevops-extensions:index:BuildFolder' 'totvsapps-dev-folder' '5687c295-d324-41c7-a430-5a50f190a0c1\/totvsapps-dev'
```

# Usage
See the example in go language [here](https://github.com/totvs/pulumi-azuredevops-extensions/tree/master/examples/go).

# Development
## Pre-requisites to develop

Install the `pulumictl` cli from the [releases](https://github.com/pulumi/pulumictl/releases) page or follow the [install instructions](https://github.com/pulumi/pulumictl#installation)

> NB: Usage of `pulumictl` is optional. If not using it, hard code the version in the [Makefile](Makefile) of when building explicitly pass version as `VERSION=0.0.1 make build`

## Build and Test

```bash
# build and install the resource provider plugin
$ make build install

# test
$ cd examples/go
$ yarn link @pulumi/azuredevops-extensions
$ yarn install
$ pulumi stack init test
$ pulumi up
```