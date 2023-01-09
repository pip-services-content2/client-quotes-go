package build

import (
	clients1 "github.com/pip-services-content2/client-quotes-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type QuotesClientFactory struct {
	*cbuild.Factory
}

func NewQuotesClientFactory() *QuotesClientFactory {
	c := &QuotesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-quotes", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-quotes", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-quotes", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewQuotesNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewQuotesMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewQuotesCommandableHttpClientV1)

	return c
}
