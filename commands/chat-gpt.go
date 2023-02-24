package Commands

import (
	"context"
	"fmt"
	"os"

	openai "jbelbo/guidoBot/internal/openai"
)

func ChatGPT() string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	ctx := context.Background()
	model := "davinci"
	prompt := "Hello, world!"
	maxTokens := 5

	text, err := client.GenerateText(ctx, model, prompt, maxTokens)
	if err != nil {
		fmt.Println("Error generating text:", err)
	}

	return text
}
