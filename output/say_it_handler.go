package output

import (
	"fmt"
	"go-say-it/inputparser"
	"net/url"
)

type Output struct {
	urls []string
}

type Handler interface {
	generateUrl(input *inputparser.Input) (*Output, error)
}

type HandlerImpl struct {
}

func (r *HandlerImpl) generateUrl(input *inputparser.Input) *Output {
	urlTemplate := "https://youglish.com/pronounce/%s/english/us"
	if len(input.Word) <= 0 {
		return nil
	}

	var urls []string
	for _, word := range input.Word {
		sayItUrl := fmt.Sprintf(urlTemplate, url.QueryEscape(word))
		urls = append(urls, sayItUrl)
	}

	return &Output{urls: urls}
}
