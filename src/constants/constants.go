// Package constants shows a pattern to group constants together
package constants

// Endpoint contains the endpoint configuration
var Endpoint struct {
	Hostname string
	Port     int
}

func init() {
	Endpoint.Hostname = "some-endpoint"
	Endpoint.Port = 9090
}
