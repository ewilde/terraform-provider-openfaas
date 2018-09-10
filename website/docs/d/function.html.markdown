---
layout: "openfaas"
page_title: "OpenFaaS: openfaas_function"
sidebar_current: "docs-openfaas-datasource-function"
description: |-
  Get information about a single OpenFaaS function.
---

# openfaas_function

Use this data source to get information about a specific function.

## Example Usage

```hcl
data "openfaas_function" "function" {
  name            = "test-function"  
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the function

## Attributes Reference

The following arugments are exported:

* `image` - (Required) Docker image that must be located in an accessible registry
* `f_process` - (Optional) Name of the watchdog process to fork
* `labels` - (Optional) Map of labels used by the back-end for making scheduling or routing decisions
* `annotations` - (Optional) Map of annotations used by the back-end for management, orchestration, events and build tasks