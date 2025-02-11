package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

type OpenAIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func queryOpenAI(apiKey, model, prompt string, maxTokens int) (string, error) {
	requestBody := OpenAIRequest{
		Model:     model,
		Messages:  []Message{{Role: "user", Content: prompt}},
		MaxTokens: maxTokens,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", err
	}

	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}
