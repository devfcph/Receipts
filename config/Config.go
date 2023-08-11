package config

var server = "localhost:9090"     // The default server address for the application.
var defaultEndpoint = "/receipts" // The default API endpoint for receipts.

// GetServer returns the server address for the application.
func GetServer() string {
	return server
}

// GetDefaultEndPoint returns the default API endpoint for receipts.
func GetDefaultEndPoint() string {
	return defaultEndpoint
}
