package inputparser

import "fmt"

type parser interface {
	parseArg(args []string) (*Input, error)
}

type ParserImpl struct {
}

func (i ParserImpl) parseArg(args []string) (*Input, error) {
	input := &Input{}
	if len(args) <= 1 {
		return nil, fmt.Errorf("Arg length is <= 1")
	}
	commandArgs := args[1:(len(args) - 1)]
	wordsToSay := make([]string, len(commandArgs))
	for _, arg := range commandArgs {
		wordsToSay = append(wordsToSay, arg)
	}
	input.word = wordsToSay
	return input, nil
}
