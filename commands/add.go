package Commands

import (
    "jbelbo/guidoBot/telegram"
)


func Add(responseBody *Telegram.MessageResponse ) error {

    responseBody.Text = "/add: Not implemented."
    return nil
}

