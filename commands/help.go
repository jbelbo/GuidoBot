package Commands

import (
	"jbelbo/guidoBot/telegram"
)

func Help(responseBody *Telegram.MessageResponse) error {

	responseBody.Text = "/help: Not implemented."
	return nil
}
