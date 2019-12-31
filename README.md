Repository which contains code and markdown for the Go patterns articles:
* [Part 1](https://dev.to/napicellatwit/golang-patterns-5a64)
* [Part 2](https://dev.to/napicellatwit/golang-patterns---part-2-1906)
* [Part 3](https://dev.to/napicella/golang-patterns-part-3-apo)

# Part 1
## Group constants
The title said it all, we want to group common constants in the same namespace.  
Code for the chain [here](https://github.com/napicella/Go-patterns/tree/master/src/constants).


## Chain
A chain of suppliers that returns as soon as one of the suppliers returns a non-zero result or an error.   
The code shows an example in which we need to load a configuration value from one of the possible sources: env variable, config file or a database.  
We want to stop searching for the config value as soon as a non-zero result is returned.  
Code for the chain [here](https://github.com/napicella/Go-patterns/tree/master/src/chain).


## Options
Options shows a flexible way to construct an object.  
The major benefit is being able to add more parameters in the future to the object constructor without breaking the clients.
In other languages, you would probably use an overloaded constructor or fall back to the builder pattern.  
Code for the option [here](https://github.com/napicella/Go-patterns/tree/master/src/options).

## Maybe
Maybe is a container which may or may not contain a non-null value.  
Code for the Maybe [here](https://github.com/napicella/Go-patterns/tree/master/src/maybe).

## Function type
Functions are first class citizen in Golang. They can be used as a type whenever we want to easily implement a Strategy pattern or similar.  
This pattern is heavily used in the golang [http package](https://golang.org/pkg/net/http/#Handler).  
Code for the Function type [here](https://github.com/napicella/Go-patterns/tree/master/src/functiontype).

# Part 2
## One liner if statement
```golang
if err := someOperation(); err != nil {
    // error handling here
}
```

I personally do not use it that much, but it can save a couple of keystrokes if all you want to do is returning the error upstream.  

```go
if err := someOperation(); err != nil {
    return errors.Wrap(err, "Ops, unable to complete some operation")
}
```

## Walker
I have seen this one used pretty much in all the codebases I stumbled upon.  
Most notably, it is used in the [path/filepath package](https://golang.org/pkg/path/filepath/#Walk).

The need is similar to the one addressed by the iterator pattern: 
* decoupling the algorithm from the data structure implementation
* separate the logic necessary to iterate over data from the one which acts on it

```go
func WalkInWordGrams(walker WordgramWalker) error {
   // Iteration happens here
}

WalkInWordGrams(func(wordGram *wordgrams.WordGram) error {
		for _, stats := range wordGram.Stats() {
			// ... do something meaningful with stats
		}

		return nil
	})
```  

In the code above we call a function `WalkInWordGrams` which takes as parameter a lambda which process data.  
The lambda func is called for each element in the data structure.  
In case of an unrecoverable error, we can stop the iteration by returning it.   

Using what we know from the previous article, we can improve our DSL by adding a type for the lamdbda function.  


```go
type WordgramWalker func(wordgram *wordgrams.WordGram) error

func WalkInWordGrams(walker WordgramWalker) error {
   // Iteration happens here
}

WalkInWordGrams(func(wordGram *wordgrams.WordGram) error {
		for _, stats := range wordGram.Stats() {
			// ... do something meaningful with stats
		}

		return nil
	})

```

## Test data
This is something pretty useful when writing tests.  
Go build ignores directory named `testdata `and when it runs tests it sets current directory as package directory. This allows you to use relative path testdata directory as a place to load and store our data.  
For example we can write an utility function like the following:  
```go
func LoadTestFile(name string, failureHandler func(message string)) []byte {
	path := filepath.Join("testdata", name)
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		failureHandler(e.Error())
	}

	return bytes
}
```

The function reads the file from testdata and returns it to the caller.  
In case it fails, it calls the failure handler.  
We can use it like this:  
```go
import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test suite")
}

var _ = Describe("Some test which requires a file", func() {
	It("does what is supposed to do", func() {
		image := LoadTestFile("amalfi-coast.jpg", Fail)
                ...
```
## Godoc Example
This is something pretty cool, which helps to improve docs for our code.  
From the golang docs:  
> Godoc examples are snippets of Go code that are displayed as package documentation and that are verified by running them as tests. They can also be run by a user visiting the godoc web page for the package and clicking the associated "Run" button.  

Writing one is really easy. We need to write a function which name starts with `Example` in a test file.  
```go
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
```
The cool part is that the example becomes a test.  
The `Output comment` is used to verify that the output matches our expectations.  
It does so, by capturing data written to standard output and then comparing the output against the example's "Output:" comment. The test passes if the test's output matches its output comment.
