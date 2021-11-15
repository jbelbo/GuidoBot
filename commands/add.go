package Commands

import (
    "fmt"
	"jbelbo/guidoBot/telegram"
)

func Add(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

    fmt.Println("Datos +v", body )
	responseBody.Text = "/add: Not implemented."
	return nil
}
