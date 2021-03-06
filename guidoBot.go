package main

import (
	Handler "jbelbo/guidoBot/handler"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/"+os.Getenv("TOKEN"), Handler.Run)
	http.ListenAndServe(":"+port, nil)
}
