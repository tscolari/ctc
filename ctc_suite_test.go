package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var CtcBin string

func TestCtc(t *testing.T) {
	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		ctcBin, err := gexec.Build("github.com/tscolari/ctc")
		Expect(err).NotTo(HaveOccurred())

		return []byte(ctcBin)
	}, func(data []byte) {
		CtcBin = string(data)
	})

	RunSpecs(t, "Ctc Suite")
}
