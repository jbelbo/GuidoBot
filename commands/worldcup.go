package Commands

import (
	"context"
	"fmt"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Match struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MatchNumber int                `bson:"MatchNumber,omitempty"`
	RoundNumber int                `bson:"RoundNumber,omitempty"`
	DateUtc     string             `bson:"DateUtc,omitempty"`
	Location    string             `bson:"Location,omitempty"`
	HomeTeam    string             `bson:"HomeTeam,omitempty"`
	AwayTeam    string             `bson:"AwayTeam,omitempty"`
	Group       string             `bson:"Group,omitempty"`
}

func MatchesForTeam(reqBody *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {
	team := strings.TrimSpace(strings.TrimPrefix(reqBody.Message.Text, "/matches"))

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("worldcup").Collection("fixture")
	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"HomeTeam", team}},
				bson.D{{"AwayTeam", team}},
			},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	var out strings.Builder
	for _, result := range results {
		out.WriteString(fmt.Sprintln(result))
		fmt.Println(result)
	}
	responseBody.Text = out.String()

	return nil

}
