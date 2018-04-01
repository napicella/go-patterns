// options shows a flexible way to construct an object.
// The major benefit is being able to add more parameters in the future to the object constructor
// without breaking the clients.
// In other languages, you would probably use an overloaded constructor or fall back
// to the builder pattern.
package options

import "fmt"

type Option func(greeting *Greeting)

// NewGreeting creates a Greeting
func NewGreeting(options ...Option) *Greeting {
	greeting := &Greeting{
		name: "Stranger",
	}

	for _, option := range options {
		option(greeting)
	}

	return greeting
}

func Name(name string) Option {
	return func(greeting *Greeting) {
		greeting.name = name
	}
}

type Greeting struct {
	name string
}

func (g *Greeting) get() string {
	return fmt.Sprintf("Hello %s", g.name)
}
