package test_helpers

import (
	"fmt"
	"os"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func BuildBoshDirector() (boshdir.Director, error) {

	logger := boshlog.NewLogger(boshlog.LevelError)
	factory := boshdir.NewFactory(logger)

	// Build a Director config from address-like string.
	// HTTPS is required and certificates are always verified.
	config, err := boshdir.NewConfigFromURL(BoshEnvironment())
	if err != nil {
		return nil, fmt.Errorf("building director config: %s", err)
	}

	// Configure custom trusted CA certificates.
	// If nothing is provided default system certificates are used.
	config.CACert = BoshCaCert()

	// Allow Director to fetch UAA tokens when necessary.
	uaa, err := buildUAA()
	if err != nil {
		return nil, fmt.Errorf("building uaa: %s", err)
	}

	config.TokenFunc = boshuaa.NewClientTokenSession(uaa).TokenFunc

	return factory.New(config, boshdir.NewNoopTaskReporter(), boshdir.NewNoopFileReporter())
}

func BoshDeployment() string {
	return os.Getenv("BOSH_DEPLOYMENT")
}

func BoshEnvironment() string {
	return os.Getenv("BOSH_ENVIRONMENT")
}

func BoshClient() string {
	return os.Getenv("BOSH_CLIENT")
}

func BoshClientSecret() string {
	return os.Getenv("BOSH_CLIENT_SECRET")
}

func BoshCaCert() string {
	return os.Getenv("BOSH_CA_CERT")
}

func buildUAA() (boshuaa.UAA, error) {
	logger := boshlog.NewLogger(boshlog.LevelError)
	factory := boshuaa.NewFactory(logger)

	// Build a UAA config from a URL.
	// HTTPS is required and certificates are always verified.

	config, err := boshuaa.NewConfigFromURL(fmt.Sprintf("https://%s:8443", BoshEnvironment()))
	if err != nil {
		return nil, fmt.Errorf("ERROR build uaa config: %s", err)
	}

	// Set client credentials for authentication.
	// Machine level access should typically use a client instead of a particular user.
	config.Client = BoshClient()
	config.ClientSecret = BoshClientSecret()

	// Configure trusted CA certificates.
	// If nothing is provided default system certificates are used.
	config.CACert = BoshCaCert()

	return factory.New(config)
}

func HostsForInstanceGroup(deployment boshdir.Deployment, instanceGroupName string) ([]string, error) {
	instances, err := deployment.Instances()
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, instance := range instances {
		if instance.Group == instanceGroupName {
			addresses = append(addresses, instance.IPs[0])
		}
	}

	return addresses, nil
}
