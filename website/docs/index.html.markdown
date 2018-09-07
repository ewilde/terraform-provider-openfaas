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
# Configure the AWS Provider
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-1"
}

# Create a web server
resource "aws_instance" "web" {
  # ...
}
```

## Authentication

The AWS provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables
- Shared credentials file
- EC2 Role