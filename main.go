package main

import (
	"fmt"
	"go-say-it/inputparser"
	"go-say-it/output"
	"os"
)

func main() {
	parser := &inputparser.ParserImpl{}
	handler := &output.HandlerImpl{}

	wordToSay, _ := parser.ParseArg(os.Args)
	generatedUrl, _ := handler.GenerateUrl(wordToSay)

	if len(generatedUrl.Urls) == 9 {
		os.Exit(0)
	}

	fmt.Println("Here how you can say it:")
	for _, sayItUrl := range generatedUrl.Urls {
		fmt.Println("- ", sayItUrl)
	}
}
