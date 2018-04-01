package functiontype

import "fmt"

func ExampleFunctionType() {
	request := Request{user: "Mickey"}

	fmt.Println(
		GreetingService(request, func(name string) string {
			return fmt.Sprintf("Hola %s!", name)
		}),
	)
	// Output: Service says: Hola Mickey!
}
