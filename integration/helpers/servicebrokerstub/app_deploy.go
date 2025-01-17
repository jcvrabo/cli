package servicebrokerstub

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/onsi/ginkgo/v2"

	"code.cloudfoundry.org/cli/v9/integration/helpers"
	. "github.com/onsi/gomega"
)

var appOrg = "fakeservicebroker"

const (
	appNamePrefix      = "hydrabroker"
	appSpace           = "integration"
	defaultMemoryLimit = "32M"
	pathToApp          = "../../assets/hydrabroker"
)

var once sync.Once
var mutex sync.Mutex

func initialize() {
	testName := os.Getenv("CF_INT_TEST_NAME")
	if len(testName) > 0 {
		appOrg += "-" + testName
	}
}

func ensureAppIsDeployed() {
	if !appResponds() {
		ensureAppIsPushed()
		Eventually(appResponds()).Should(BeTrue())
	}
}

func appResponds() bool {
	resp, err := http.Head(appURL())
	Expect(err).ToNot(HaveOccurred())
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNoContent
}

func ensureAppIsPushed() {
	appExists := func() bool {
		session := helpers.CF("app", "--guid", appName())
		session.Wait()
		return session.ExitCode() == 0
	}

	pushApp := func() bool {
		session := helpers.CF(
			"push", appName(),
			"-p", pathToApp,
			"-m", defaultMemoryLimit,
			"-b", "https://github.com/cloudfoundry/go-buildpack.git", // Some legacy envs have buildpack that's too old
		)
		session.Wait()
		return session.ExitCode() == 0
	}

	cleanupAppsFromPreviousRuns := func() {
		session := helpers.CF("apps").Wait()

		if session.ExitCode() == 0 {
			matchingApps := regexp.MustCompile(fmt.Sprintf(`%s-\d+`, appNamePrefix)).
				FindAllString(string(session.Out.Contents()), -1)

			for _, app := range matchingApps {
				if app != appName() {
					session := helpers.CF("delete", app, "-f")
					session.Wait()
				}
			}
		}
	}

	once.Do(func() {
		initialize()
	})

	// mutex protects from goroutines, and we retry later to protect from other test processes
	mutex.Lock()
	defer mutex.Unlock()

	helpers.CreateOrgAndSpaceUnlessExists(appOrg, appSpace)
	helpers.WithRandomHomeDir(func() {
		helpers.SetAPI()
		helpers.LoginCF()
		helpers.TargetOrgAndSpace(appOrg, appSpace)

		cleanupAppsFromPreviousRuns()

		ok := false
		for attempts := 0; attempts < 5 && !ok; attempts++ {
			ok = appExists()
			if !ok {
				ok = pushApp()
			}
			if !ok {
				time.Sleep(time.Second)
			}
		}

		Expect(ok).To(BeTrue(), "Failed to push app")
	})
}

func appURL(paths ...string) string {
	return fmt.Sprintf("http://%s.%s%s", appName(), helpers.DefaultSharedDomain(), strings.Join(paths, ""))
}

func appName() string {
	id := ginkgo.GinkgoRandomSeed()
	if len(os.Getenv("REUSE_SERVICE_BROKER_APP")) > 0 {
		id = 0
	}
	return fmt.Sprintf("%s-%010d", appNamePrefix, id)
}
