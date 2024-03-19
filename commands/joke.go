package Commands

import (
	"encoding/json"
	"fmt"
	"io"
	Telegram "jbelbo/guidoBot/telegram"
	"net/http"

	"github.com/rs/zerolog/log"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error while reading Open Joke response")

		return err
	}

	log.Debug().Msg(`GetJoke: ` + string(body))

	var joke JokeResponse
	err = json.Unmarshal(body, &joke)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling open joke response")

		return err
	}

	responseBody.Text = joke.Url + "\n" + joke.Value

	return nil
}
