package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPruebasHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PruebasHttp Suite")
}
