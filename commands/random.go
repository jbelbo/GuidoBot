package Commands

import (
	"context"
	"fmt"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Fields struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Text string             `bson:"txt,omitempty"`
}

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

	pipeline := []bson.D{{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}}
	showInfoCursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showsWithInfo []Fields
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showsWithInfo[0].Text)

	responseBody.Text = showsWithInfo[0].Text

	return nil

}

func RandomStuffWithKeyword(body *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("messages").Collection("originals")

	pipeline := []bson.D{{{Key: "$match", Value: bson.D{{Key: "txt", Value: bson.M{"$regex": "crypto|BTC|ETH|XTZ|Tezos", "$options": "im"}}}}}, {{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}}
	showInfoCursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showsWithInfo []Fields
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showsWithInfo[0].Text)

	responseBody.Text = showsWithInfo[0].Text

	return nil

}
