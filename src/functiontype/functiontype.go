// functiontype shows that functions are first class citizen in Golang.
// They can be used as a type whenever we want to easily implement a Strategy pattern or similar.
// This pattern is heavily used in the golang http package https://golang.org/pkg/net/http/#Handler
package functiontype

import "fmt"

type Greeting func(name string) string

func GreetingService(request Request, greeting Greeting) string {
	return fmt.Sprintf("Service says: %s", greeting(request.user))
}

type Request struct {
	user string
}
