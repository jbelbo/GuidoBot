package Commands

import (
	Telegram "jbelbo/guidoBot/telegram"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func Send(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	text := strings.TrimPrefix(body.Message.Text, "/send ")
	stringSlice := strings.Split(text, "=")
	chatID, err := strconv.ParseInt(stringSlice[0], 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("Error while parsing chatID")
		return err
	}

	responseBody.ChatID = chatID
	responseBody.Text = stringSlice[1]

	return nil
}
