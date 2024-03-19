package Commands

import Telegram "jbelbo/guidoBot/telegram"

// TODO: implement status command that describes the current status of all the services provided by the bot
func Status(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	responseBody.Text = "Status: active"

	return nil
}
