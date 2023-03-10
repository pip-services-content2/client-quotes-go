package version1

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type QuoteV1 struct {
	Id      string            `json:"id"`
	Text    map[string]string `json:"text"`
	Author  map[string]string `json:"author"`
	Status  string            `json:"status"`
	Tags    []string          `json:"tags"`
	AllTags []string          `json:"all_tags"`
}

func NewQuoteV1(text, author map[string]string, status string, tags, allTags []string) *QuoteV1 {
	c := &QuoteV1{
		Id:     data.IdGenerator.NextLong(),
		Text:   text,
		Author: author,
	}

	if status == "" {
		c.Status = QuoteStatusNew
	} else {
		c.Status = status
	}

	if tags == nil {
		c.Tags = make([]string, 0)
	} else {
		c.Tags = tags
	}

	if allTags == nil {
		c.AllTags = make([]string, 0)
	} else {
		c.AllTags = allTags
	}

	return c
}
