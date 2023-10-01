package output

import (
	"fmt"
	"go-say-it/inputparser"
	"net/url"
)

type Output struct {
	Urls []string
}

type Handler interface {
	GenerateUrl(input *inputparser.Input) (*Output, error)
}

type HandlerImpl struct {
}

func (r *HandlerImpl) GenerateUrl(input *inputparser.Input) (*Output, error) {
	urlTemplate := "https://youglish.com/pronounce/%s/english/us"
	if len(input.Word) <= 0 {
		return nil, fmt.Errorf("empty input")
	}

	var urls []string
	for _, word := range input.Word {
		sayItUrl := fmt.Sprintf(urlTemplate, url.QueryEscape(word))
		urls = append(urls, sayItUrl)
	}

	return &Output{Urls: urls}, nil
}
