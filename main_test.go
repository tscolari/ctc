package main_test

import (
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("filters out logs that are not inside the time range", func() {
		cmd := exec.Command(CtcBin, "--time", "2017-05-17T15:00:00", "--interval", "5s", "assets/test.log")
		output := gbytes.NewBuffer()
		sess, err := gexec.Start(cmd, output, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(sess).Should(gexec.Exit(0))
		entries := strings.Split(string(output.Contents()), "\n")
		Expect(entries).To(HaveLen(6))
	})
})
