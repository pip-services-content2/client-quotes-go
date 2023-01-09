package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-quotes-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type quotesCommandableHttpClientV1Test struct {
	client  *version1.QuotesCommandableHttpClientV1
	fixture *QuotesClientFixtureV1
}

func newQuotesCommandableHttpClientV1Test() *quotesCommandableHttpClientV1Test {
	return &quotesCommandableHttpClientV1Test{}
}

func (c *quotesCommandableHttpClientV1Test) setup(t *testing.T) *QuotesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewQuotesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewQuotesClientFixtureV1(c.client)

	return c.fixture
}

func (c *quotesCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newQuotesCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
