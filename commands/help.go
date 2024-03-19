package Commands

import Telegram "jbelbo/guidoBot/telegram"

// TODO: implement help command that describes the available commands and their usage
func Help(responseBody *Telegram.MessageResponse) error {
	responseBody.Text = "/help: Not implemented."

	return nil
}
