package main

import (
	"fmt"
	"os"
	"strings"
)

var apiKeyEnvs = []struct {
	env      string
	provider *LLMProvider
}{
	{"OPENAI_API_KEY", openai},
	{"GEMINI_API_KEY", gemini},
	{"DEEPSEEK_API_KEY", deepseek},
}

func main() {
	cmd := strings.Join(os.Args[1:], " ")
	if cmd == "" {
		fmt.Println("usage: howto <query>")
		os.Exit(1)
	}

	provider, apiKey := detectProvider()
	if provider == nil {
		fmt.Fprintln(os.Stderr, "error: no API key found")
		fmt.Fprintln(os.Stderr, "set one of: OPENAI_API_KEY, GEMINI_API_KEY, DEEPSEEK_API_KEY")
		os.Exit(1)
	}

	model := os.Getenv("HOWTO_MODEL")

	response, err := queryLLM(provider, apiKey, model, prompt(cmd))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	insertInput(sanitizeCommand(response))
}

func detectProvider() (*LLMProvider, string) {
	for _, e := range apiKeyEnvs {
		if key := os.Getenv(e.env); key != "" {
			return e.provider, key
		}
	}
	return nil, ""
}
