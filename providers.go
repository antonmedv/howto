package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LLMProvider struct {
	Name         string
	Endpoint     string
	DefaultModel string
}

var (
	openai = &LLMProvider{
		Name:         "OpenAI",
		Endpoint:     "https://api.openai.com/v1/chat/completions",
		DefaultModel: "gpt-4o",
	}
	gemini = &LLMProvider{
		Name:         "Gemini",
		Endpoint:     "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions",
		DefaultModel: "gemini-2.0-flash",
	}
	deepseek = &LLMProvider{
		Name:         "DeepSeek",
		Endpoint:     "https://api.deepseek.com/chat/completions",
		DefaultModel: "deepseek-chat",
	}
)

type ChatRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error *APIError `json:"error,omitempty"`
}

type APIError struct {
	Message string `json:"message"`
}

func queryLLM(provider *LLMProvider, apiKey, model, prompt string) (string, error) {
	if model == "" {
		model = provider.DefaultModel
	}

	requestBody := ChatRequest{
		Model:     model,
		Messages:  []Message{{Role: "user", Content: prompt}},
		MaxTokens: 1000,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", provider.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", err
	}

	if chatResp.Error != nil {
		return "", fmt.Errorf("API error: %s", chatResp.Error.Message)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from %s", provider.Name)
}
