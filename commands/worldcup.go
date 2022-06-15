package Commands

import (
	"context"
	"fmt"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"os"
	"regexp"
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

type Flag struct {
	Name  string `bson:"name,omitempty"`
	Emoji string `bson:"emoji,omitempty"`
}

func LookupFlag(country string) string {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("worldcup").Collection("flags")
	fmt.Print(country)
	filter := bson.D{{"name", bson.D{{"$regex", regexp.QuoteMeta(country)}, {"$options", "i"}}}}

	var result Flag
	if err = collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return ""
	}

	return result.Emoji
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
				bson.D{{"HomeTeam", bson.D{{"$regex", regexp.QuoteMeta(team)}, {"$options", "i"}}}},
				bson.D{{"AwayTeam", bson.D{{"$regex", regexp.QuoteMeta(team)}, {"$options", "i"}}}},
			},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var results []Match
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	var out strings.Builder
	for _, result := range results {
		out.WriteString(fmt.Sprintf("On %s: %s(%s) vs. %s(%s) in %s\n", result.DateUtc, result.HomeTeam, LookupFlag(result.HomeTeam), result.AwayTeam, LookupFlag(result.AwayTeam), result.Location))
	}
	responseBody.Text = out.String()

	return nil

}
