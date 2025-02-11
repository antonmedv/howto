package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cmd := strings.Join(os.Args[1:], " ")
	if cmd == "" {
		fmt.Println("usage: howto <query>")
		os.Exit(1)
	}

	openAIKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		panic("OPENAI_API_KEY not set")
	}

	openAiModel := "gpt-4o"
	model, ok := os.LookupEnv("OPENAI_MODEL")
	if ok {
		openAiModel = model
	}

	response, err := queryOpenAI(openAIKey, openAiModel, prompt(cmd), 1000)
	if err != nil {
		panic(err)
	}

	insertInput(sanitizeCommand(response))
}
