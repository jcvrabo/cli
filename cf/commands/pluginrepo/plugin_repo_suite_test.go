package pluginrepo_test

import (
	"code.cloudfoundry.org/cli/v9/cf/commands/pluginrepo"
	"code.cloudfoundry.org/cli/v9/cf/i18n"
	"code.cloudfoundry.org/cli/v9/cf/util/testhelpers/configuration"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPluginRepo(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = pluginrepo.RepoPlugins{}

	RegisterFailHandler(Fail)
	RunSpecs(t, "PluginRepo Suite")
}
