package main

import (
	_ "github.com/lib/pq"
	"jbelbo/guidoBot/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.ListenAndServe(":"+port, http.HandlerFunc(Handler.Run))
}
