# Elastic Path Commerce Cloud Terraform Provider

This repository contains a [Terraform](https://www.terraform.io) provider for the [Elastic Path Commerce Cloud API](https://documentation.elasticpath.com/commerce-cloud/docs/api/). 

This project was based upon the [Terraform Provider Scaffolding Project](https://github.com/hashicorp/terraform-provider-scaffolding)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.15.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Environment Variables

The following environment variables are [defined in Account Management](internal/config/env.go), and can be used to influence behaviour.

| Name                      | Default          | Description                                                                                                                                        |
| --------------------------| ---------------- | ------------ |
| `EPCC_API_BASE_URL`       | -                | The Base URL of the EPCC API                                          |
| `EPCC_BETA_API_FEATURES`  | -                | The value of the `EP-Beta-Features` header being sent to the EPCC API |
| `EPCC_CLIENT_ID`          | -                | Client ID used for authenticating to the EPCC API                     |
| `EPCC_CLIENT_SECRET`      | -                | Client Secret used for authenticating to the EPCC API                 |
| `EPCC_LOG_DIR`            | (work directory) | The directory of the http log files                                   |

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.


## Project Layout

This project follows the terraform provider template: [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)

| Directory         | Description                                                                                    |
|-------------------|------------------------------------------------------------------------------------------------|
| `docs/`            | Folder that contains documentation                                                            |
| `examples/`        | Directory for sample resources and data sources                                               |
| `component-tests/` | Component tests for the service are located in here.                                          |
| `external/`        | Any Go package that can be shared with other projects                                         |
| `internal/`        | Application specific Go packages, e.g., they cannot be shared and are specific to this service|

## Using the provider

You would use the epcc-terraform-provider just as any other terraform provider. See the `./examples` directory for sample resources and data sources.

See the [Core Terraform Workflow] (https://www.terraform.io/guides/core-workflow.html) page for more info on using Terraform.

## Useful commands

| Command         | Description                                                                                   |
| ----------------| ----------------------------------------------------------------------------------------------|
| go install      | Compile the provider. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.|
| go generate     | Generate or update documentation                                                                       |
| make testacc    | Runs the full suite of Acceptance tests                                             |

## Debugging the Provider

1. Run `make build`
   
2. Run the following command (assuming you've installed delve)
```bash
dlv exec --headless ./bin/terraform-provider-my-provider -- --debug
```

3. Connect with your Debugger
   
4. Find the line `TF_REATTACH_PROVIDERS` in the output

5. When running terraform prefix the above to the command, for example:

```
TF_REATTACH_PROVIDERS='...' terraform apply
```

[Debugging Providers](https://www.terraform.io/docs/extend/debugging.html#starting-a-provider-in-debug-mode)
### Useful Links

1. [AWS SDK for Go](https://github.com/aws/aws-sdk-go-v2)
2. [Extending Terraform](https://www.terraform.io/docs/extend/index.html)
3. [AWS Terraform Provider](https://github.com/hashicorp/terraform-provider-aws)