---
layout: "openfaas"
page_title: "OpenFaaS: openfaas_function"
sidebar_current: "docs-openfaas-resource-function"
description: |-
  Provides an OpenFaaS function resource.
---

# openfaas_function

Provides an OpenFaaS function resource.

## Example Usage

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

## Argument Reference

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

### Function Resources Arguments 

The meanings and formats of `limits` and `requests` may vary depending on whether you are using Kubernetes or Docker Swarm. In general:

 - Reserve maintains the host resources to ensure that the container can use them
 - Limits specify the maximum amount of host resources that a container can consume

See docs for [Docker Swarm](https://docs.docker.com/config/containers/resource_constraints/) or for [Kubernetes](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#how-pods-with-resource-limits-are-run).

* `cpu` - (Optional) Amount of cpu required
* `memory` - (Optional) Amount of memory required
