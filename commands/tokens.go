package Commands

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"jbelbo/guidoBot/telegram"
	"net/http"
	"strings"
)

type Token struct {
	Price  float32 `json:"current_price"`
	ID     string
	Symbol string
	Name   string
	Ath    float32
	Atl    float32
	Vol    int64 `json:"total_volume"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func formatAllTokens(responseBody *Telegram.MessageResponse, tokens []Token) {
	responseBody.Text = "= Lista: \n"
	for _, value := range tokens {
		responseBody.Text = responseBody.Text + value.Name + "(" + value.Symbol + ") = " + fmt.Sprintf("%f", value.Price) + "\n"
	}
}

func formatListedTokens(responseBody *Telegram.MessageResponse, tokens []Token, params string) {
	coins := strings.Split(strings.ToUpper(params), " ")
	var requestedTokens []Token

	for _, value := range tokens {
		if contains(coins, strings.ToUpper(value.Symbol)) {
			requestedTokens = append(requestedTokens, value)
		}
	}
	if len(requestedTokens) == 0 {
		return
	}

	responseBody.Text = "= Detalles: \n"
	for _, token := range requestedTokens {
		responseBody.Text = "*" + responseBody.Text + token.Name + "(" + token.Symbol + ")* :\n"
		responseBody.Text = responseBody.Text + "    - Price: " + fmt.Sprintf("%f", token.Price) + "\n"
		responseBody.Text = responseBody.Text + "    - ATL: " + fmt.Sprintf("%f", token.Atl) + " - ATH: " + fmt.Sprintf("%f", token.Ath) + "\n"
		responseBody.Text = responseBody.Text + "    - VOL: " + fmt.Sprintf("%d", token.Vol) + "\n"
	}
}

// List Token from API
//
func ListTokens(reqBody *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	//https://www.sohamkamani.com/golang/json/
	//https://dev.to/billylkc/parse-json-api-response-in-go-10ng
	params := strings.TrimSpace(strings.TrimPrefix(reqBody.Message.Text, "/tokens "))

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//fmt.Println(string(body))

	if err == nil {
		var tokens []Token
		json.Unmarshal([]byte(body), &tokens)
		//fmt.Printf("Tokens : %+v", tokens)

		if len(params) == 0 {
			formatAllTokens(responseBody, tokens)
		} else {
			formatListedTokens(responseBody, tokens, params)
		}
	}

	return nil
}
