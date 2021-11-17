package Commands

import (
	"context"
	"database/sql"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RandomStuff(responseBody *Telegram.MessageResponse) error {

	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("messages").Collection("originals")

	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		return err
	}
	id := res.InsertedID.(primitive.ObjectID).String()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT message FROM phrase ORDER BY random() LIMIT 1")
	if err != nil {
		log.Fatal("Error while querying DB")
	}

	defer results.Close()

	for results.Next() {
		var err = results.Scan(&responseBody.Text)
		responseBody.Text = responseBody.Text + " " + id
		if err != nil {
			log.Fatal("Error while reading from row")
		}
		return nil
	}

	return nil
}
