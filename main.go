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

	wordToSay, err := parser.ParseArg(os.Args)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	generatedUrl, err := handler.GenerateUrl(wordToSay)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	if len(generatedUrl.Urls) == 9 {
		os.Exit(0)
	}

	fmt.Println("Here how you can say it:")
	for _, sayItUrl := range generatedUrl.Urls {
		fmt.Println("- ", sayItUrl)
	}
}
