---
layout: "openfaas"
page_title: "Provider: OpenFaaS"
sidebar_current: "docs-openfaas-index"
description: |-
  The OpenFaaS provider is used to interact OpenFaaS functions. 
  The provider may need to be configured with the proper credentials if using basic auth.
---

# OpenFaaS Provider

The OpenFaaS provider is used to interact with OpenFaaS functions.
The provider may need to be configured with the proper credentials to talk to the OpenFaaS
 gateway if using basic auth (recommended). 


Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the OpenFaaS Provider
provider "openfaas" {
  uri       = "https://localhost:8080"
  user_name = "a-username"
  password  = "a-password"
}

# Create a function
resource "openfaas_function" "figlet" {
  # ...
}
```

## Authentication

The OpenFaaS provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Basic authentication
- No authentication


### Basic authentication
This is the *recommended* option, configure the provider block inline:

```hcl
provider "openfaas" {
  uri       = "https://localhost:8080"
  user_name = "a-username"
  password  = "a-password"
}
```

### No authentication
This is *not* recommended, please only use in a private environment:

```hcl
provider "openfaas" {
  uri       = "http://localhost:8080"  
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the AWS
 `provider` block:

* `uri` - (Optional) This [OpenFaaS gateway](https://docs.openfaas.com/conceptual/#api-gateway-ui-portal) uri. 
If omitted, default value is `http://localhost:8080`.

* `tls_insecure` - (Optional) Explicitly allow the provider to perform "insecure" SSL requests. 
If omitted, default value is `false`.

* `user_name` - (Optional) This is the OpenFaaS [basic authentication user name](https://docs.openfaas.com/reference/authentication/#for-the-api-gateway).
If ommited, basic authentication is not used.

* `password` - (Optional) This is the OpenFaaS [basic authentication passwrd](https://docs.openfaas.com/reference/authentication/#for-the-api-gateway).
If ommited, basic authentication is not used.
