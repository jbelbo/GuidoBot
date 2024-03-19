package Commands

import (
	"context"
	Telegram "jbelbo/guidoBot/telegram"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Add(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	text := strings.TrimPrefix(body.Message.Text, "/add ")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error().Err(err).Msg("Error while connecting to mongo")

		return err
	}

	collection := client.Database("messages").Collection("custom")
	res, err := collection.InsertOne(context.Background(), bson.M{"txt": text})
	if err != nil {
		log.Error().Err(err).Msg("Error while inserting to mongo")

		return err
	}

	id := res.InsertedID.(primitive.ObjectID).String()
	responseBody.Text = "/add: " + id

	return nil
}
