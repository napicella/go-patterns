package chain

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler")
}

var _ = Describe("testing chain of handlers", func() {

	Describe("chain stops as soon as a non empty value is returned", func() {

		Context("first handler in the chain returns non empty string", func() {
			It("returns the value returned by the first handler", func() {
				Expect(
					chain(
						func() (string, error) {
							return "FIRST", nil
						}, func() (string, error) {
							return "SECOND", nil
						},
					).get(),
				).To(Equal("FIRST"))
			})
		})

		Context("nth handler in the chain return non empty string", func() {
			It("returns the value returned by the nth handler", func() {
				suppliers := createSuppliers(5)
				suppliers = append(suppliers, func() (string, error) {
					return "NTH", nil
				})

				suppliers = append(suppliers, createSuppliers(5)...)

				Expect(chain(suppliers...).get()).To(Equal("NTH"))
			})
		})

		Context("last handler in the chain return non empty string", func() {
			It("returns the value returned by the last handler", func() {
				suppliers := createSuppliers(10)
				suppliers = append(suppliers, func() (string, error) {
					return "LAST", nil
				})

				Expect(chain(suppliers...).get()).To(Equal("LAST"))
			})
		})
	})

	Describe("chain stops as soon as one of the handler returns an error", func() {

		Context("nth handler in the chain return an error", func() {
			It("returns the error returned by the nth handler", func() {
				suppliers := createSuppliers(5)
				suppliers = append(suppliers, func() (string, error) {
					return "", errors.New("OPS...something wrong")
				})

				called := false
				suppliers = append(suppliers, func() (string, error) {
					called = true
					return "", nil
				})

				_, e := chain(suppliers...).get()

				Expect(e).To(MatchError(errors.New("OPS...something wrong")))
				Expect(called).To(BeFalse())
			})
		})
	})

	Describe("edge cases", func() {

		Context("an empty chain", func() {
			It("returns an empty string", func() {
				Expect(chain().get()).To(BeEmpty())
			})
		})
	})
})

func createSuppliers(n int) []func() (string, error) {
	var suppliers []func() (string, error)

	for i := 0; i < n; i++ {
		suppliers = append(
			suppliers,
			func() (string, error) {
				return "", nil
			})
	}

	return suppliers
}
