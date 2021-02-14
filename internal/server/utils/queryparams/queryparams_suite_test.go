package queryparams_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueryparams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queryparams Suite")
}
