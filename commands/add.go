package Commands

import (
	"jbelbo/guidoBot/telegram"
)

func Add(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	responseBody.Text = "/add: Not implemented."
	return nil
}
