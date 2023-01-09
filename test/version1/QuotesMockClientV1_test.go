package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-quotes-go/version1"
)

type quotesMockClientV1Test struct {
	client  *version1.QuotesMockClientV1
	fixture *QuotesClientFixtureV1
}

func newQuotesMockClientV1Test() *quotesMockClientV1Test {
	return &quotesMockClientV1Test{}
}

func (c *quotesMockClientV1Test) setup(t *testing.T) *QuotesClientFixtureV1 {
	c.client = version1.NewQuotesMockClientV1()

	c.fixture = NewQuotesClientFixtureV1(c.client)

	return c.fixture
}

func (c *quotesMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockOperations(t *testing.T) {
	c := newQuotesMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
