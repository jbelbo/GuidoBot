package Telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Send a response according to the environment.
func SendResponse(chatID int64, message *MessageResponse) error {

	// Create the JSON body from the struct
	responseBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	//HEROKU
	heroku := os.Getenv("HEROKU")
	envHeroku, _ := strconv.ParseBool(heroku)

	if envHeroku == true {
		apiKey := os.Getenv("API_KEY")
		res, err := http.Post("https://api.telegram.org/bot"+apiKey+"/sendMessage", "application/json", bytes.NewBuffer(responseBytes))
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return errors.New("unexpected status" + res.Status)
		}
	} else {
		fmt.Println("Response is ", message)
	}
	return nil
}
