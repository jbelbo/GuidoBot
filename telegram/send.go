package Telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

// Sends a response according to the environment.
func SendResponse(chatID int64, message *MessageResponse) error {
	responseBytes, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("Error while marshalling response")

		return err
	}

	heroku := os.Getenv("HEROKU")
	envHeroku, err := strconv.ParseBool(heroku)
	if err != nil {
		log.Error().Err(err).Msg("Error while parsing HEROKU env var")

		return err
	}

	if envHeroku {
		apiKey := os.Getenv("API_KEY")
		res, err := http.Post("https://api.telegram.org/bot"+apiKey+"/sendMessage", "application/json", bytes.NewBuffer(responseBytes))
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return errors.New("unexpected status" + res.Status)
		}
	} else {
		log.Debug().Msg("Response is: " + message.Text)
	}
	return nil
}
