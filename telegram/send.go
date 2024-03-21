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
		return err
	}

	devModeEnvVar := os.Getenv("DEV_MODE")
	devMode, err := strconv.ParseBool(devModeEnvVar)
	if err != nil {
		return err
	}

	if devMode {
		log.Debug().Msg("Response is: " + message.Text)

		return nil
	}

	apiKey := os.Getenv("API_KEY")
	res, err := http.Post("https://api.telegram.org/bot"+apiKey+"/sendMessage", "application/json", bytes.NewBuffer(responseBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + res.Status)
	}

	return nil
}
