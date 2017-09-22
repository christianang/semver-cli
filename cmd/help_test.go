package cmd_test

import (
	"bytes"

	"github.com/christianang/semver-cli/cmd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Help", func() {
	Describe("Execute", func() {
		var (
			stdout *bytes.Buffer
			stderr *bytes.Buffer

			command cmd.Help
		)

		BeforeEach(func() {
			stdout = bytes.NewBuffer([]byte{})
			stderr = bytes.NewBuffer([]byte{})

			command = cmd.NewHelp(stdout, stderr)
		})

		It("prints usage", func() {
			err := command.Execute([]string{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stderr.String()).To(ContainSubstring("Usage: semver"))
		})
	})
})
