package main

import (
	"fmt"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Compare", func(version1, version2, expectedReturn string) {
	stdout, exitCode := semver("compare", version1, version2)
	Expect(exitCode).To(Equal(0))
	Expect(stdout).To(Equal(fmt.Sprintf("%s\n", expectedReturn)))
},
	Entry("when both semvers are equal", "1.11.1", "1.11.1", "0"),
	Entry("when semver one is greater than semver two (major)", "2.11.1", "1.11.1", "1"),
	Entry("when semver one is less than semver two (major)", "1.11.1", "2.11.1", "-1"),
	Entry("when semver one is greater than semver two (minor)", "1.12.1", "1.11.1", "1"),
	Entry("when semver one is less than semver two (minor)", "1.11.1", "1.12.1", "-1"),
	Entry("when semver one is greater than semver two (patch)", "1.11.2", "1.11.1", "1"),
	Entry("when semver one is less than semver two (patch)", "1.11.1", "1.11.2", "-1"),
)
