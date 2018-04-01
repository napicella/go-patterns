package chain

import "fmt"

func ExampleChain() {
	endpoint, _ := chain(
		loadEndpointFromConfigFile,
		loadEndpointFromEnvVariables,
		loadEndpointFromDatabase,
	).get()

	fmt.Println(endpoint)
	// Output: some-endpoint
}

func loadEndpointFromEnvVariables() (string, error) {
	return "", nil
}

func loadEndpointFromConfigFile() (string, error) {
	return "", nil
}

func loadEndpointFromDatabase() (string, error) {
	return "some-endpoint", nil
}
