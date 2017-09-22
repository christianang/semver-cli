package cmd_test

import (
	"bytes"
	"errors"

	"github.com/blang/semver"
	"github.com/christianang/semver-cli/cmd"
	"github.com/christianang/semver-cli/cmd/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate counterfeiter -o ./fakes/semver.go --fake-name Semver . _semver
type _semver interface {
	Make(s string) (semver.Version, error)
	Compare(v1, v2 semver.Version) int
}

var _ = Describe("Compare", func() {
	Describe("Execute", func() {
		var (
			stdout *bytes.Buffer
			stderr *bytes.Buffer

			fakeSemver *fakes.Semver

			command cmd.Compare
		)

		BeforeEach(func() {
			stdout = bytes.NewBuffer([]byte{})
			stderr = bytes.NewBuffer([]byte{})

			fakeSemver = &fakes.Semver{}
			cmd.SetSemverMake(fakeSemver.Make)
			cmd.SetSemverCompare(fakeSemver.Compare)

			fakeSemver.CompareReturns(9999)

			command = cmd.NewCompare(stdout, stderr)
		})

		AfterEach(func() {
			cmd.ResetSemverMake()
		})

		It("prints the return value of a semver compare", func() {
			err := command.Execute([]string{"1.11.5", "1.12.5"})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeSemver.MakeArgsForCall(0)).To(Equal("1.11.5"))
			Expect(fakeSemver.MakeArgsForCall(1)).To(Equal("1.12.5"))
			Expect(stdout.String()).To(Equal("9999\n"))
		})

		Context("when an error occurs", func() {
			Context("when the first semver is invalid", func() {
				It("returns an error", func() {
					fakeSemver.MakeReturnsOnCall(0, semver.Version{}, errors.New("failed to make"))
					err := command.Execute([]string{"1.x.x.x", "1.12.5"})
					Expect(err).To(MatchError("failed to make"))
				})
			})

			Context("when the first semver is invalid", func() {
				It("returns an error", func() {
					fakeSemver.MakeReturnsOnCall(1, semver.Version{}, errors.New("failed to make"))
					err := command.Execute([]string{"1.x.x.x", "1.12.5"})
					Expect(err).To(MatchError("failed to make"))
				})
			})
		})
	})
})
