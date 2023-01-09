package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IQuotesClientV1 interface {
	GetQuotes(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*QuoteV1], error)

	GetRandomQuote(ctx context.Context, correlationId string, filter *data.FilterParams) (*QuoteV1, error)

	GetQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error)

	CreateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error)

	UpdateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error)

	DeleteQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error)
}
