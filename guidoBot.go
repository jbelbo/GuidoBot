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

// https://core.telegram.org/bots/api#sendmessage
type messageResponse struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}


// Decode and Parse
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}

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


// Random command
//
// TODO: randomize message selection
func randomStuff(responseBody *messageResponse ) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT message FROM phrase LIMIT 1")
	if err != nil {
		log.Fatal("Error while querying DB")
	}

	defer results.Close()

	for results.Next() {
		var err = results.Scan(&responseBody.Text)
		if err != nil {
			log.Fatal("Error while reading from row")
        }
        return nil
	}

    return nil
}



//
// Available commands::
//
//    /help
//    /add
//    /random
//    hola = /random
//
func parseRequest(body *webhookReqBody) error {

	// Create the request body struct
	responseBody := messageResponse{
		ChatID: body.Message.Chat.ID,
		Text:   "",
	}

    //Process hola command
    if strings.HasPrefix(strings.ToLower(body.Message.Text), "hola") {
        var err = randomStuff(&responseBody)
        if err != nil {
            log.Fatal("Error retriving random stuff")
        }
	}


    //Process /random command
    if strings.HasPrefix(strings.ToLower(body.Message.Text), "/random") {
        var err = randomStuff(&responseBody)
        if err != nil {
            log.Fatal("Error retriving random stuff")
        }
	}

    //Process /help command
    if strings.HasPrefix(strings.ToLower(body.Message.Text), "/add") {
        responseBody.Text = "/add: Not implemented."
	}


    //Process /help command
    if strings.HasPrefix(strings.ToLower(body.Message.Text), "/help") {
        responseBody.Text = "/help: Not implemented."
	}

    if responseBody.Text == "" {
        return nil
    }

    return sendResponse(body.Message.Chat.ID, &responseBody)
}


// Send a response according to the environment.
func sendResponse(chatID int64, message *messageResponse) error {

	// Create the JSON body from the struct
	responseBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}


    //HEROKU
    if heroku:= os.Getenv("HEROKU"); heroku == "true" {
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
    return nil;
}

func main() {

    port := os.Getenv("PORT");
    if port == "" {
		log.Fatal("$PORT must be set")
    }

	http.ListenAndServe(":"+port, http.HandlerFunc(Handler))
}
