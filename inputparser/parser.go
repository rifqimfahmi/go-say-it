package inputparser

import "fmt"

type parser interface {
	ParseArg(args []string) (*Input, error)
}

type ParserImpl struct {
}

func (i ParserImpl) ParseArg(args []string) (*Input, error) {
	if len(args) <= 1 {
		return nil, fmt.Errorf("arg length is <= 1")
	}
	var wordsToSay []string
	for _, arg := range args[1:] {
		wordsToSay = append(wordsToSay, arg)
	}
	input := &Input{Word: wordsToSay}
	return input, nil
}
