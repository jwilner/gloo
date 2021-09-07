package translator_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	namespace = "translator-test-ns"

	_ = BeforeSuite(func() {
		err := os.Setenv(statusutils.PodNamespaceEnvName, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

	_ = AfterSuite(func() {
		err := os.Unsetenv(statusutils.PodNamespaceEnvName)
		Expect(err).NotTo(HaveOccurred())
	})
)

func TestTranslator(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Translator Suite", []Reporter{junitReporter})
}
