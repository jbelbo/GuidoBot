package Commands

import (
	"context"
	"os"
	"strconv"
	"strings"

	openai "jbelbo/guidoBot/internal/openai"
	Telegram "jbelbo/guidoBot/telegram"
)

func ChatGPT(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	ctx := context.Background()
	model := os.Getenv("OPENAI_API_MODEL")
	prompt := strings.TrimPrefix(body.Message.Text, "/chatgpt ")
	maxTokens, err := strconv.ParseInt(os.Getenv("OPENAI_API_MAX_TOKENS"), 10, 64)
	if err != nil {
		return err
	}

	text, err := client.GenerateText(ctx, model, prompt, int(maxTokens))
	if err != nil {
		return err
	}

	responseBody.Text = text

	return nil
}
