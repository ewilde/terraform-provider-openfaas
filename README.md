The terraform provider for [OpenFaaS](https://www.openfaas.com/)

## Documentation

Full documentation, see: https://openfaas-tfe.edwardwilde.com/docs/providers/openfaas

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

[![image](https://user-images.githubusercontent.com/329397/45926773-920cbd80-bf1f-11e8-9b26-88dc5df0fc7e.png)](https://www.youtube.com/watch?v=sSctTy6YIlU&feature=youtu.be)

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

## Installing the provider
Download the latest [release](https://github.com/ewilde/terraform-provider-openfaas/releases/latest) or build the provider from source, then follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

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
