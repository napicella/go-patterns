package options

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestOption(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Options")
}

var _ = Describe("testing Maybe", func() {

	Context("Greeting with no Name option", func() {
		It("returns default greeting", func() {
			greeting := NewGreeting()
			Expect(greeting.get()).To(Equal("Hello Stranger"))
		})
	})

	Context("Greeting with Name option", func() {
		It("returns custom greeting", func() {
			greeting := NewGreeting(Name("Mickey"))
			Expect(greeting.get()).To(Equal("Hello Mickey"))
		})
	})
})
