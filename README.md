# Terraform provider for OpenFaaS

[![Build Status](https://travis-ci.org/ewilde/terraform-provider-openfaas.svg?branch=master)](https://travis-ci.org/ewilde/terraform-provider-openfaas)

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

## Building and Installing

Since this isn't maintained by Hashicorp, you have to install it manually. There
are two main ways:

### Download a release

Download and unzip the [latest
release](https://github.com/ewilde/terraform-provider-openfaas/releases/latest).

Then, move the binary to your terraform plugins directory. [The
docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins)
don't fully describe where this is.

* On Mac, it's `~/.terraform.d/plugins/darwin_amd64`
* On Linux, it's `~/.terraform.d/plugins/linux_amd64`
* On Windows, it's `$APPDATA\terraform.d\plugins\windows_amd64`

### Build using the Makefile

Install [Go](https://www.golang.org/) v1.9+ on your machine; clone the source,
and let `make install` do the rest.

#### Mac

```sh
brew install go  # or upgrade
brew install dep # or upgrade
mkdir -p $GOPATH/src/github.com/ewilde; cd $GOPATH/src/github.com/ewilde
git clone https://github.com/ewilde/terraform-provider-openfaas 
cd terraform-provider-openfaas
make install
# it may take a while to download `hashicorp/terraform`. be patient.
```

#### Linux

Install go and dep from your favourite package manager or from source. Then:

```sh
mkdir -p $GOPATH/src/github.com/ewilde; cd $GOPATH/src/github.com/ewilde
git clone https://github.com/ewilde/terraform-provider-openfaas 
cd terraform-provider-openfaas
make install
# it may take a while to download `hashicorp/terraform`. be patient.
```

#### Windows

In PowerShell, running as Administrator:

```powershell
choco install golang
choco install dep
choco install zip
choco install git # for git-bash
choco install make
```

In a shell that has Make, like Git-Bash:

```sh
mkdir -p $GOPATH/src/github.com/ewilde; cd $GOPATH/src/github.com/ewilde
git clone https://github.com/ewilde/terraform-provider-openfaas 
cd terraform-provider-openfaas
make install
# it may take a while to download `hashicorp/terraform`. be patient.
```


## Developing the Provider

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

> Note: At the moment the acceptance tests assume OpenFaaS gateway is running on http://localhost:8080 *without* 
basic authentication enabled.

```sh
$ make testacc
```

## Building the documentation
Currently a bit manual ¯\_(ツ)_/¯

1. build the content
```sh
$ make website-build
```

2. copy build output into `gh-pages` branch of this repo
```sh
$ git checkout gh-pages 
cp -r ../terraform-website/content/build/docs/providers/openfaas/ ./docs/providers/openfaas
```