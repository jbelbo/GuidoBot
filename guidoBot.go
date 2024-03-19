package main

import (
	Handler "jbelbo/guidoBot/handler"
	logger "jbelbo/guidoBot/internal/log"
	"jbelbo/guidoBot/internal/utils"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	logLevel := utils.GetEnv("LOG_LEVEL", "INFO")
	logMode := utils.GetEnv("LOG_MODE", "PRETTY")
	logger.InitLogger()
	logger.SetLoggingLevel(logLevel)
	logger.SetLoggingMode(logMode)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal().Msg("$PORT must be set")
	}

	http.HandleFunc("/"+os.Getenv("TOKEN"), Handler.Run)
	http.ListenAndServe(":"+port, nil)
}
