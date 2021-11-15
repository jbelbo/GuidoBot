package Commands

import (
	"database/sql"
	"log"
	"os"
    "jbelbo/guidoBot/telegram"
)


func RandomStuff(responseBody *Telegram.MessageResponse ) error {
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

