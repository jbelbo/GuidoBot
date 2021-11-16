package Commands

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"jbelbo/guidoBot/telegram"
	"net/http"
)

type Token struct {
	Price  float32 `json:"current_price"`
	ID     string
	Symbol string
	Name   string
}

// List Token from API
//
func ListTokens(responseBody *Telegram.MessageResponse) error {

	//https://www.sohamkamani.com/golang/json/
	//https://dev.to/billylkc/parse-json-api-response-in-go-10ng

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
		fmt.Printf("Tokens : %+v", tokens)
		responseBody.Text = "= Lista: \n"

		for _, value := range tokens {
			responseBody.Text = responseBody.Text + value.Name + "(" + value.Symbol + ") = " + fmt.Sprintf("%f", value.Price) + "\n"
		}
	}

	return nil
}
