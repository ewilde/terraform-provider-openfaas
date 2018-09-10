The terra

## openfaas_function

Provides an OpenFaaS function resource.

### Example Usage

```hcl
resource "openfaas_function" "function_test" {
  name            = "test-function"
  image           = "functions/alpine:latest"
  f_process       = "sha512sum"
  labels {
    Group       = "London"
    Environment = "Test"
  }

  annotations {
    CreatedDate = "Mon Sep  3 07:15:55 BST 2018"
  }
}
```

### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the function
* `image` - (Required) Docker image that must be located in an accessible registry
* `network` - (Optional) Name of the provider network to associate this function with
* `f_process` - (Optional) Name of the watchdog process to fork
* `env_vars` - (Optional) Map of environment variables for the function to use
* `registry_auth` - (Optional) Docker private registry authentication string which is base64-encoded (as present in ~/.docker/config.json)
* `constraints` - (Optional) List of deployment constraints, which are specific to the configured OpenFaaS Provider. i.e. "node.platform.os == linux"
* `secrets` - (Optional) List of names of secrets that are required to be loaded from the Docker Swarm
* `labels` - (Optional) Map of labels used by the back-end for making scheduling or routing decisions
* `annotations` - (Optional) Map of annotations used by the back-end for management, orchestration, events and build tasks
* `limits` - (Optional) [Function resources block](#function-resources-arguments) used to configure maximum amount of resources.
* `requests` - (Optional) [Function resources block](#function-resources-arguments) used to configure minimum amount of resources.

#### Function Resources Arguments 

The meanings and formats of `limits` and `requests` may vary depending on whether you are using Kubernetes or Docker Swarm. In general:

 - Reserve maintains the host resources to ensure that the container can use them
 - Limits specify the maximum amount of host resources that a container can consume

See docs for [Docker Swarm](https://docs.docker.com/config/containers/resource_constraints/) or for [Kubernetes](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#how-pods-with-resource-limits-are-run).

* `cpu` - (Optional) Amount of cpu required
* `memory` - (Optional) Amount of memory required


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/ewilde/terraform-provider-openfaas`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:ewilde/terraform-provider-openfaas
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/ewilde/terraform-provider-openfaas
$ make build
```

## Using the provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-openfaas
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```