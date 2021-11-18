package Commands

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jbelbo/guidoBot/telegram"
	"log"
	"os"
	"strings"
	"time"
)

func Add(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	text := strings.TrimPrefix(body.Message.Text, "/add ")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("messages").Collection("custom")

	res, err := collection.InsertOne(context.Background(), bson.M{"txt": text})
	if err != nil {
		return err
	}
	id := res.InsertedID.(primitive.ObjectID).String()

	responseBody.Text = "/add: " + id
	return nil
}
