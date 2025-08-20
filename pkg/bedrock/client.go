package bedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	bedrockruntime "github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Client struct {
	br *bedrockruntime.Client
}

// ClaudeRequest represents the request structure for Claude model
type ClaudeRequest struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature,omitempty"`
	StopSequences     []string `json:"stop_sequences,omitempty"`
}

// ClaudeResponse represents the response structure from Claude model
type ClaudeResponse struct {
	Completion string `json:"completion"`
}

func New(region string) *Client {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("AWS config error: %v", err)
	}
	return &Client{br: bedrockruntime.NewFromConfig(cfg)}
}

func (c *Client) Ask(prompt string, maxTokens int) (string, error) {
	ctx := context.Background()
	modelID := "anthropic.claude-instant-v1"

	// Format prompt according to Claude requirements
	enclosedPrompt := "Human: " + prompt + "\n\nAssistant:"

	request := ClaudeRequest{
		Prompt:            enclosedPrompt,
		MaxTokensToSample: maxTokens,
		Temperature:       0.5,
		StopSequences:     []string{"\n\nHuman:"},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.br.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(modelID),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
		Body:        body,
	})
	if err != nil {
		return "", fmt.Errorf("bedrock api call failed: %w", err)
	}

	var response ClaudeResponse
	if err := json.Unmarshal(resp.Body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return response.Completion, nil
}
