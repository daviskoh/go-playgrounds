package variadic_funcs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVariadicFuncs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VariadicFuncs Suite")
}
