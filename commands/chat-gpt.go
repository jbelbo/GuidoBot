package Commands

import (
	"context"
	"os"
	"strings"

	openai "jbelbo/guidoBot/internal/openai"
	Telegram "jbelbo/guidoBot/telegram"
)

func ChatGPT(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	ctx := context.Background()
	model := "davinci"
	prompt := strings.TrimPrefix(body.Message.Text, "/chatgpt ")
	maxTokens := 5

	text, err := client.GenerateText(ctx, model, prompt, maxTokens)
	if err != nil {
		return err
	}

	responseBody.Text = text

	return nil
}
