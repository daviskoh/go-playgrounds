package variadic_funcs_test

import (
	. "github.com/daviskoh/go-playgrounds/variadic_funcs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VariadicFuncs", func() {
	It("returns arg when only 1 arg specified", func() {
		Expect(Add(1)).To(Equal(1))
	})

	It("adds 2 nums", func() {
		Expect(Add(3, 4)).To(Equal(7))
	})

	It("adds multiple nums together", func() {
		Expect(Add(1, 2, 3, 4, 5)).To(Equal(1 + 2 + 3 + 4 + 5))
	})
})
