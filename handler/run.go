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

func Run(res http.ResponseWriter, req *http.Request) {
	body := &Telegram.WebhookReqBody{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("Could not decode request body", err)
		return
	}

	if err := parseRequest(body); err != nil {
		fmt.Println("Error while sending reply: ", err)
		return
	}

	fmt.Println("Reply sent")
}

// Available commands:: /help /add /random /tokens /send /status @user_mention
func parseRequest(body *Telegram.WebhookReqBody) error {
	responseBody := Telegram.MessageResponse{
		ChatID: body.Message.Chat.ID,
		Text:   "",
	}

	regex := regexp.MustCompile("^\\/[a-zA-Z]*")
	command := regex.FindString(body.Message.Text)

	if command != "" {
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
			err = Commands.Send(body, &responseBody)
		case "/help":
			err = Commands.Help(&responseBody)
		case "/weather":
			err = Commands.GetWeather(body, &responseBody)
		case "/joke":
			err = Commands.GetJoke(body, &responseBody)
		case "/matches":
			err = Commands.MatchesForTeam(body, &responseBody)
		case "/chatgpt":
			err = Commands.ChatGPT(body, &responseBody)
		default:
			err = nil
		}

		if err != nil {
			log.Fatal("Error while processing command")
		}

		return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
	}

	if recognizeKeyword(body) {
		var err = Commands.RandomStuffWithKeyword(body, &responseBody)
		if err != nil {
			log.Fatal("Error in random quote with keyword command")
		}

		return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
	}

	if messageContainsMention(body.Message.Entities) {
		var err = Commands.RandomStuff(&responseBody)
		if err != nil {
			log.Fatal("Error in mention command")
		}

		return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
	}

	if responseBody.Text == "" {
		return nil
	}

	return Telegram.SendResponse(body.Message.Chat.ID, &responseBody)
}

//messageContainsMention this method recognizes when a user has been mentioned
func messageContainsMention(entities []Telegram.MessageEntity) bool {
	for _, entity := range entities {
		if entity.Type == "mention" || entity.Type == "text_mention" {
			return true
		}
	}
	return false
}

func recognizeKeyword(body *Telegram.WebhookReqBody) bool {
	matched, _ := regexp.MatchString(`(?i)crypto|ETH|BTC`, body.Message.Text)

	return matched
}
