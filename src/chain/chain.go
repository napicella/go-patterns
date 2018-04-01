// chain shows a chain of suppliers that returns as soon as one of the suppliers
// returns a non-zero result or an error
package chain

// chain builds a chain of responsibility using the provided suppliers as handlers:
// the suppliers are wrapped in shortCircuitOnErrorHandler, which means the chain stops as soon as one of the suppliers
// returns an error.
// Returns the first not empty string or an empty string in case all suppliers return an empty string.
// The suppliers are arranged in the order provided in the argument.
func chain(suppliers ...func() (string, error)) handler {
	if len(suppliers) == 0 {
		return &nilHandler{}
	}

	first, suppliers := suppliers[0], suppliers[1:]

	return &shortCircuitOnErrorHandler{
		supplier: first,
		next:     chain(suppliers...),
	}
}

// handler interface for a chain of responsibility in which handlers return either a string or an error
type handler interface {
	get() (string, error)
}

type nilHandler struct{}

func (h *nilHandler) get() (string, error) {
	return "", nil
}

// shortCircuitOnErrorHandler returns an error as soon as an error is encountered, otherwise calls the next in the chain
// until a non empty string is returned or the end of the chain is reached
type shortCircuitOnErrorHandler struct {
	next     handler
	supplier func() (string, error)
}

func (b *shortCircuitOnErrorHandler) get() (string, error) {
	result, e := b.supplier()

	if e != nil {
		return "", e
	}

	if result != "" {
		return result, nil
	}

	return b.next.get()
}
