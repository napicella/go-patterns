package maybe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMaybe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maybe")
}

var _ = Describe("testing Maybe", func() {

	Context("User present", func() {
		var greeting string
		MaybeUser(getUser(1)).IfPresent(func(u *User) {
			greeting = "Hello " + u.name
		})

		It("greets the user", func() {
			Expect(greeting).To(Equal("Hello Mickey"))
		})
	})

	Context("User absent", func() {
		var greeting string
		MaybeUser(getUser(-1)).WhenAbsent(func() {
			greeting = "Hello stranger"
		})

		It("greets the user", func() {
			Expect(greeting).To(Equal("Hello stranger"))
		})
	})

})
