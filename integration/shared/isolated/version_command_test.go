package isolated

import (
	"strings"

	"code.cloudfoundry.org/cli/v9/integration/helpers"

	"github.com/blang/semver/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("version command", func() {
	DescribeTable("displays version",
		func(arg string) {
			session := helpers.CF(arg)
			Eventually(session).Should(Exit(0))
			output := string(session.Out.Contents())
			version := strings.Split(output, " ")[2]
			versionNumber := strings.Split(version, "+")[0]
			_, err := semver.Make(versionNumber)
			Expect(err).To(Not(HaveOccurred()))
			Eventually(session).ShouldNot(Say("cf version 0.0.0-unknown-version"))
		},

		Entry("when passed version", "version"),
		Entry("when passed -v", "-v"),
		Entry("when passed --version", "--version"),
	)
})
