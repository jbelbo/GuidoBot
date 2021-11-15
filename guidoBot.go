package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
    "jbelbo/guidoBot/handler"
)



func main() {

    port := os.Getenv("PORT");
    if port == "" {
		log.Fatal("$PORT must be set")
    }

	http.ListenAndServe(":"+port, http.HandlerFunc(Handler.Run))
}
