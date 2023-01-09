package version1

import (
	"context"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/random"
)

type QuotesMockClientV1 struct {
	quotes []*QuoteV1
}

func NewQuotesMockClientV1() *QuotesMockClientV1 {
	return &QuotesMockClientV1{
		quotes: make([]*QuoteV1, 0),
	}
}

func (c *QuotesMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[i] {
				return true
			}
		}
	}

	return false
}

func (c *QuotesMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *QuotesMockClientV1) matchMultilanguageString(value map[string]string, search string) bool {
	for _, text := range value {
		if c.matchString(text, search) {
			return true
		}
	}

	return false
}

func (c *QuotesMockClientV1) matchSearch(item *QuoteV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchMultilanguageString(item.Text, search) {
		return true
	}
	if c.matchMultilanguageString(item.Author, search) {
		return true
	}
	if c.matchString(item.Status, search) {
		return true
	}
	return false
}

func (c *QuotesMockClientV1) composeFilter(filter *data.FilterParams) func(*QuoteV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	status, statusOk := filter.GetAsNullableString("name")
	author, authorOk := filter.GetAsNullableString("author")

	tagsString := filter.GetAsString("tags")
	tags := make([]string, 0)

	exceptIdsString := filter.GetAsString("except_ids")
	exceptIds := make([]string, 0)

	// Process tags filter
	if tagsString != "" {
		tags = data.TagsProcessor.CompressTags([]string{tagsString})
	}

	// Process except ids filter
	if exceptIdsString != "" {
		exceptIds = strings.Split(exceptIdsString, ",")
	}

	return func(item *QuoteV1) bool {
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		if idOk && item.Id != id {
			return false
		}
		if len(exceptIds) > 0 && c.contains(exceptIds, []string{item.Id}) {
			return false
		}
		if authorOk && c.matchMultilanguageString(item.Author, author) {
			return false
		}
		if statusOk && item.Status != status {
			return false
		}
		if len(tags) > 0 && !c.contains(item.AllTags, tags) {
			return false
		}
		return true
	}
}

func (c *QuotesMockClientV1) GetQuotes(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*QuoteV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*QuoteV1, 0)
	for _, v := range c.quotes {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}
	return *data.NewDataPage(items, len(c.quotes)), nil
}

func (c *QuotesMockClientV1) GetRandomQuote(ctx context.Context, correlationId string, filter *data.FilterParams) (*QuoteV1, error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*QuoteV1, 0)
	for _, v := range c.quotes {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}

	buf := *items[random.Integer.Next(0, len(items))]
	return &buf, nil
}

func (c *QuotesMockClientV1) GetQuoteById(ctx context.Context, correlationId string, quoteId string) (result *QuoteV1, err error) {
	for _, v := range c.quotes {
		if v.Id == quoteId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *QuotesMockClientV1) CreateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	if quote.Status == "" {
		quote.Status = QuoteStatusNew
	}
	if quote.Tags == nil {
		quote.Tags = make([]string, 0)
	}
	if quote.AllTags == nil {
		quote.Tags = data.TagsProcessor.ExtractHashTags("#text#author")
	}

	buf := *quote
	c.quotes = append(c.quotes, &buf)

	return quote, nil
}

func (c *QuotesMockClientV1) UpdateQuote(ctx context.Context, correlationId string, quote *QuoteV1) (*QuoteV1, error) {
	if quote == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.quotes {
		if v.Id == quote.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	if quote.Status == "" {
		quote.Status = QuoteStatusNew
	}
	if quote.Tags == nil {
		quote.Tags = make([]string, 0)
	}
	if quote.AllTags == nil {
		quote.Tags = data.TagsProcessor.ExtractHashTags("#text#author")
	}

	buf := *quote
	c.quotes[index] = &buf
	return quote, nil
}

func (c *QuotesMockClientV1) DeleteQuoteById(ctx context.Context, correlationId string, quoteId string) (*QuoteV1, error) {
	var index = -1
	for i, v := range c.quotes {
		if v.Id == quoteId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.quotes[index]
	if index < len(c.quotes) {
		c.quotes = append(c.quotes[:index], c.quotes[index+1:]...)
	} else {
		c.quotes = c.quotes[:index]
	}
	return item, nil
}
