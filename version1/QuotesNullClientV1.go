package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type QuoteNullClientV1 struct {
}

func NewQuotesNullClientV1() *QuoteNullClientV1 {
	return &QuoteNullClientV1{}
}

func (c *QuoteNullClientV1) GetQuotes(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*QuoteV1], error) {
	return *data.NewEmptyDataPage[*QuoteV1](), nil
}

func (c *QuoteNullClientV1) GetRandomQuote(ctx context.Context, correlationId string, filter *data.FilterParams) (*QuoteV1, error) {
	return nil, nil
}

func (c *QuoteNullClientV1) GetQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error) {
	return nil, nil
}

func (c *QuoteNullClientV1) CreateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	return quote, nil
}

func (c *QuoteNullClientV1) UpdateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	return quote, nil
}

func (c *QuoteNullClientV1) DeleteQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error) {
	return nil, nil
}
