package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type QuotesCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewQuotesCommandableHttpClientV1() *QuotesCommandableHttpClientV1 {
	return &QuotesCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/quotes"),
	}
}

func (c *QuotesCommandableHttpClientV1) GetQuotes(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*QuoteV1], error) {
	res, err := c.CallCommand(ctx, "get_quotes", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *data.NewEmptyDataPage[*QuoteV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*QuoteV1]](res, correlationId)
}

func (c *QuotesCommandableHttpClientV1) GetRandomQuote(ctx context.Context, correlationId string, filter *data.FilterParams) (*QuoteV1, error) {
	res, err := c.CallCommand(ctx, "get_random_quote", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*QuoteV1](res, correlationId)
}

func (c *QuotesCommandableHttpClientV1) GetQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error) {
	res, err := c.CallCommand(ctx, "get_quote_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"quote_id", quoteId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*QuoteV1](res, correlationId)
}

func (c *QuotesCommandableHttpClientV1) CreateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	res, err := c.CallCommand(ctx, "create_quote", correlationId, data.NewAnyValueMapFromTuples(
		"quote", quote,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*QuoteV1](res, correlationId)
}

func (c *QuotesCommandableHttpClientV1) UpdateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	res, err := c.CallCommand(ctx, "update_quote", correlationId, data.NewAnyValueMapFromTuples(
		"quote", quote,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*QuoteV1](res, correlationId)
}

func (c *QuotesCommandableHttpClientV1) DeleteQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error) {
	res, err := c.CallCommand(ctx, "delete_quote_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"quote_id", quoteId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*QuoteV1](res, correlationId)
}
