package config

var server = "localhost:9090"
var defaultEndpoint = "/receipts"

func GetServer() string {
	return server
}

func GetDefaultEndPoint() string {
	return defaultEndpoint
}
