package inputparser

import (
	"reflect"
	"testing"
)

func TestParserValidCmdArg(t *testing.T) {
	command := []string{"bin", "say this", "and that"}
	parser := ParserImpl{}

	input, _ := parser.parseArg(command)

	if len(input.word) != 2 {
		t.Error("Input length expected to be 2, but got", len(input.word), ". content: ", input.word)
	}

	if !reflect.DeepEqual(input.word, []string{"say this", "and that"}) {
		t.Error("unexpected content:", input.word)
	}
}
