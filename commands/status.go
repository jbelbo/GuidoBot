package Commands

import (
	"jbelbo/guidoBot/telegram"
)

func Status(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	responseBody.Text = "Status: active"

	return nil
}
