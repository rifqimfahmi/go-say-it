package output

import (
	"go-say-it/inputparser"
	"testing"
)

func TestSayItHandler(t *testing.T) {
	input := &inputparser.Input{
		Word: []string{"Hi there", "What about this"},
	}
	handler := &HandlerImpl{}

	words, _ := handler.GenerateUrl(input)

	if words.Urls[0] != "https://youglish.com/pronounce/Hi+there/english/us" {
		t.Error("unexpected output: ", words.Urls[0])
	}
}
