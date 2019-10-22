package errback_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestErrback(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errback Suite")
}
