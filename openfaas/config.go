package openfaas

type Config struct {
	TLSInsecure bool
	GatewayURI  string
	GatewayUserName string
	GatewayPassword string
}
