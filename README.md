# Permissions


## Environment

You need to set Azure DevOps token and project with:

```sh
export AZDO_PERSONAL_ACCESS_TOKEN=ahdflakjsdhdfkajsdhdflkajshdflakjs34t9813h4v134ht3i4
export AZDO_ORG_SERVICE_URL=https://dev.azure.com/typeorganizationhere
```
or
```sh
pulumi config set --secret azuredevops-extensions:config:orgServiceUrl https://dev.azure.com/typeorganizationhere
pulumi config set --secret azuredevops-extensions:config:personalAccessToken ahdflakjsdhdfkajsdhdflkajshdflakjs34t9813h4v134ht3i4
```

# Usage

Install Plugin

```bash
wget https://github.com/totvs/pulumi-azuredevops-extensions/releases/download/v0.0.2-alpha.1636747691%2Bbb5b83c8/pulumi-resource-azuredevops-extensions-v0.0.2-alpha.1636747691+bb5b83c8-linux-amd64.tar.gz

pulumi plugin install resource azuredevops-extensions 0.0.2 -f pulumi-resource-azuredevops-extensions-v0.0.2-alpha.1636747691+bb5b83c8-linux-amd64.tar.gz
```

## Pre-requisites to develop

Install the `pulumictl` cli from the [releases](https://github.com/pulumi/pulumictl/releases) page or follow the [install instructions](https://github.com/pulumi/pulumictl#installation)

> NB: Usage of `pulumictl` is optional. If not using it, hard code the version in the [Makefile](Makefile) of when building explicitly pass version as `VERSION=0.0.1 make build`

## Build and Test

```bash
# build and install the resource provider plugin
$ make build install

# test
$ cd examples/ts
$ yarn link @pulumi/azuredevops-extensions
$ yarn install
$ pulumi stack init test
$ pulumi up
```