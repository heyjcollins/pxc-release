package connection_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	helpers "specs/test_helpers"
)

func TestConnection(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PXC Acceptance Tests -- Connection")
}

var _ = BeforeSuite(func() {
	requiredEnvs := []string{
		"MYSQL_HOST",
		"MYSQL_USERNAME",
		"MYSQL_PASSWORD",
	}
	helpers.CheckForRequiredEnvVars(requiredEnvs)
})
