package Commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"net/http"
)

type JokeResponse struct {
	Url   string `json:"url"`
	Value string `json:"value"`
}

func GetJoke(reqBody *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error while reading Open Joke response")
	}
	fmt.Println(string(body))

	var joke JokeResponse
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Println("error unmarshalling open weather response: ", err)
	}
	responseBody.Text = joke.Url + "\n" + joke.Value

	return nil
}
