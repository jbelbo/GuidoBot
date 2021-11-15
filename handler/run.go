package Handler

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"jbelbo/guidoBot/commands"
	"jbelbo/guidoBot/telegram"
	"log"
	"net/http"
	"strings"
)

// Decode and Parse
func Run(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &Telegram.WebhookReqBody{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if err := parseRequest(body); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	fmt.Println("reply sent")
}

//
// Available commands::
//
//    /help
//    /add
//    /random
//    hola = /random
//
func parseRequest(body *Telegram.WebhookReqBody) error {

	// Create the request body struct
	responseBody := Telegram.MessageResponse{
		ChatID: body.Message.Chat.ID,
		Text:   "",
	}

	//Process hola command
	if strings.HasPrefix(strings.ToLower(body.Message.Text), "hola") {
		var err = Commands.RandomStuff(&responseBody)
		if err != nil {
			log.Fatal("Error in hola command")
		}
	}

	//Process /random command
	if strings.HasPrefix(strings.ToLower(body.Message.Text), "/random") {
		var err = Commands.RandomStuff(&responseBody)
		if err != nil {
			log.Fatal("Error in /random command")
		}
	}

	//Process /tokens command
	if strings.HasPrefix(strings.ToLower(body.Message.Text), "/tokens") {
		var err = Commands.ListTokens(&responseBody)
		if err != nil {
			log.Fatal("Error in /tokens command")
		}
	}

	//Process /add command
	if strings.HasPrefix(strings.ToLower(body.Message.Text), "/add") {
		var err = Commands.Add(&responseBody)
		if err != nil {
			log.Fatal("Error in /add command")
		}

	}

	//Process /help command
	if strings.HasPrefix(strings.ToLower(body.Message.Text), "/help") {
		var err = Commands.Help(&responseBody)
		if err != nil {
			log.Fatal("Error in /help command")
		}
	}

	if responseBody.Text == "" {
		return nil
	}

	return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
}
