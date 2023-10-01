package inputparser

import (
	"reflect"
	"testing"
)

func TestParserValidCmdArg(t *testing.T) {
	command := []string{"bin", "say this", "and that"}
	parser := ParserImpl{}

	input, _ := parser.ParseArg(command)

	if len(input.Word) != 2 {
		t.Error("Input length expected to be 2, but got", len(input.Word), ". content: ", input.Word)
	}

	if !reflect.DeepEqual(input.Word, []string{"say this", "and that"}) {
		t.Error("unexpected content:", input.Word)
	}
}
