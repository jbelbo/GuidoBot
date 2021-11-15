package Commands

import (
    "fmt"
	"jbelbo/guidoBot/telegram"
)

func Status(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

    responseBody.Text = fmt.Sprintf("ChatID %d", body.Message.Chat.ID) + "\n Message: " + body.Message.Text

	return nil
}
