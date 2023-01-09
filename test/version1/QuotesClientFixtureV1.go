package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-quotes-go/version1"
	"github.com/stretchr/testify/assert"
)

type QuotesClientFixtureV1 struct {
	Client version1.IQuotesClientV1
	QUOTE1 *version1.QuoteV1
	QUOTE2 *version1.QuoteV1
}

func NewQuotesClientFixtureV1(client version1.IQuotesClientV1) *QuotesClientFixtureV1 {
	return &QuotesClientFixtureV1{
		Client: client,
		QUOTE1: &version1.QuoteV1{
			Id:      "1",
			Text:    map[string]string{"en": "Text 1"},
			Author:  map[string]string{"en": "Author 1"},
			Status:  version1.QuoteStatusCompleted,
			Tags:    make([]string, 0),
			AllTags: make([]string, 0),
		},

		QUOTE2: &version1.QuoteV1{
			Id:      "2",
			Text:    map[string]string{"en": "Text 2"},
			Author:  map[string]string{"en": "Author 2"},
			Status:  version1.QuoteStatusCompleted,
			Tags:    []string{"TAG 1"},
			AllTags: []string{"tag1"},
		},
	}
}

func (c *QuotesClientFixtureV1) clear() {
	page, _ := c.Client.GetQuotes(context.Background(), "", nil, nil)

	for _, v := range page.Data {
		quote := v
		c.Client.DeleteQuoteById(context.Background(), "", quote.Id)
	}
}

func (c *QuotesClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one quote
	quote1, err := c.Client.CreateQuote(context.Background(), "", c.QUOTE1)
	assert.Nil(t, err)

	assert.NotNil(t, quote1)
	assert.Equal(t, quote1.Id, c.QUOTE1.Id)
	assert.Equal(t, quote1.Text["en"], c.QUOTE1.Text["en"])
	assert.Equal(t, quote1.Author["en"], c.QUOTE1.Author["en"])

	// Create another quote
	quote2, err := c.Client.CreateQuote(context.Background(), "", c.QUOTE2)
	assert.Nil(t, err)

	assert.NotNil(t, quote2)
	assert.Equal(t, quote2.Id, c.QUOTE2.Id)
	assert.Equal(t, quote2.Text["en"], c.QUOTE2.Text["en"])
	assert.Equal(t, quote2.Author["en"], c.QUOTE2.Author["en"])

	// Get all quotes
	page, err1 := c.Client.GetQuotes(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Update the quote
	quote1.Text["en"] = "Updated Content 1"
	quote, err := c.Client.UpdateQuote(context.Background(), "", quote1)
	assert.Nil(t, err)

	assert.NotNil(t, quote)
	assert.Equal(t, quote.Text["en"], "Updated Content 1")
	assert.Equal(t, quote.Author["en"], c.QUOTE1.Author["en"])

	// Delete quote
	_, err = c.Client.DeleteQuoteById(context.Background(), "", quote1.Id)
	assert.Nil(t, err)

	// Try to get deleted quote
	quote, err = c.Client.GetQuoteById(context.Background(), "", quote1.Id)
	assert.Nil(t, err)

	assert.Nil(t, quote)
}
