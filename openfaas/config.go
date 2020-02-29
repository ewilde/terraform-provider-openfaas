package openfaas

import "github.com/openfaas/faas-cli/proxy"

type Config struct {
	Client            *proxy.Client
	FunctionNamespace string
}
