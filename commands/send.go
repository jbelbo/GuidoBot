package Commands

import (
	"jbelbo/guidoBot/telegram"
	"strconv"
	"strings"
)

func Send(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	text := strings.TrimPrefix(body.Message.Text, "/send ")

	stringSlice := strings.Split(text, "=")
	chatID, _ := strconv.ParseInt(stringSlice[0], 10, 64)
	responseBody.ChatID = chatID

	responseBody.Text = stringSlice[1]

	return nil
}
