package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strings"
)

// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if !strings.Contains(strings.ToLower(body.Message.Text), "hola") {
		return
	}

	if err := sendResponse(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	fmt.Println("reply sent")
}

// https://core.telegram.org/bots/api#sendmessage
type messageResponse struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendResponse(chatID int64) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// ToDo randomize message selection
	results, err := db.Query("SELECT message FROM phrase")
	if err != nil {
		log.Fatal("Error while querying DB")
	}

	defer results.Close()

	// Create the request body struct
	responseBody := &messageResponse{
		ChatID: chatID,
		Text:   "",
	}

	for results.Next() {
		var err = results.Scan(&responseBody.Text)
		if err != nil {
			log.Fatal("Error while reading from row")
		}
	}


	// Create the JSON body from the struct
	responseBytes, err := json.Marshal(responseBody)
	if err != nil {
		return err
	}

	// ToDo move hardcoded token to secrets / env variables
	apiKey := os.Getenv("API_KEY")
	res, err := http.Post("https://api.telegram.org/bot"+apiKey+"/sendMessage", "application/json", bytes.NewBuffer(responseBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.ListenAndServe(":"+port, http.HandlerFunc(Handler))
}
