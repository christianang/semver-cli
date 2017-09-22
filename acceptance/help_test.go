package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Help", func() {
	It("prints useful help information", func() {
		_, stderr, exitCode := semver("help")
		Expect(exitCode).To(Equal(0))
		Expect(stderr).To(ContainSubstring(`Usage:`))
	})
})
