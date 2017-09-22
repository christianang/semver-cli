package main_test

import (
	"os/exec"
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance")
}

var pathToSemverCLI string

var _ = BeforeSuite(func() {
	var err error

	pathToSemverCLI, err = gexec.Build("github.com/christianang/semver-cli/semver")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func semver(args ...string) (string, string, int) {
	cmd := exec.Command(pathToSemverCLI, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session).Should(gexec.Exit())

	return string(session.Out.Contents()), string(session.Err.Contents()), session.ExitCode()
}
