package block

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

type HtmlScraperBlock struct {
	BaseBlock
}

func NewHtmlScraperBlock(id string, nc *nats.Conn) *HtmlScraperBlock {
	return &HtmlScraperBlock{
		BaseBlock: BaseBlock{
			ID:         id,
			Type:       "html-scraper",
			NatsClient: nc,
		},
	}
}

func (h *HtmlScraperBlock) Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error) {
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidInput
	}

	html, ok := inputMap["html"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	selector, ok := inputMap["selector"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var result []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		result = append(result, s.Text())
	})

	return result, nil
}
