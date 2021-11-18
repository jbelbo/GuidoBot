package Handler

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"jbelbo/guidoBot/commands"
	"jbelbo/guidoBot/telegram"
	"log"
	"net/http"
	"regexp"
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

// Available commands:: /help /add /random /tokens /send /status @user_mention
func parseRequest(body *Telegram.WebhookReqBody) error {

	responseBody := Telegram.MessageResponse{
		ChatID: body.Message.Chat.ID,
		Text:   "",
	}

	if botHasBeenMentioned(body.Message.Entities) {
		var err = Commands.RandomStuff(&responseBody)
		if err != nil {
			log.Fatal("Error in mention command")
		}
	}

	regex := regexp.MustCompile("^\\/[a-zA-Z]*")
	command := regex.FindString(body.Message.Text)
	var err error

	switch command {
	case "/random":
		err = Commands.RandomStuff(&responseBody)
	case "/tokens":
		err = Commands.ListTokens(body, &responseBody)
	case "/add":
		err = Commands.Add(body, &responseBody)
	case "/status":
		err = Commands.Status(body, &responseBody)
	case "/send":
		err = Commands.Help(&responseBody)
	case "/help":
		err = Commands.Help(&responseBody)
	default:
		err = nil
	}

	if err != nil {
		log.Fatal("Error in command")
	}

	if responseBody.Text == "" {
		return nil
	}

	return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
}

// ToDo it actually recognizes if any user has been mentioned, need to fix this
//botHasBeenMentioned this method recognizes when the bot has been mentioned
func botHasBeenMentioned(entities []Telegram.MessageEntity) bool {
	for _, entity := range entities {
		if entity.Type == "mention" {
			return true
		}
	}
	return false
}
