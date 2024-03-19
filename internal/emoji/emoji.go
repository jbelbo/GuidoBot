package Emoji

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"regexp"
)

type Emoji struct {
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Category    string   `json:"category"`
	Aliases     []string `json:"aliases"`
	Tags        []string `json:"tags"`
}

var emojis []Emoji = nil

func Init() error {
	if emojis != nil {
		return nil
	}

	jsonFile, err := os.Open("assets/emoji.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &emojis)

	return nil
}

func Search(word string) (string, error) {
	re, _ := regexp.Compile("(?i)" + regexp.QuoteMeta(word))
	for _, emoji := range emojis {
		if re.MatchString(emoji.Description) {
			return emoji.Emoji, nil
		}

		for _, alias := range emoji.Aliases {
			if re.MatchString(alias) {
				return emoji.Emoji, nil
			}
		}

		for _, tag := range emoji.Tags {
			if re.MatchString(tag) {
				return emoji.Emoji, nil
			}
		}
	}

	return "", errors.New("no matching emoji")
}

func SearchInCategory(category string, word string) (string, error) {
	re, _ := regexp.Compile("(?i)" + regexp.QuoteMeta(word))
	for _, emoji := range emojis {
		if emoji.Category != category {
			continue
		}

		if re.MatchString(emoji.Description) {
			return emoji.Emoji, nil
		}

		for _, alias := range emoji.Aliases {
			if re.MatchString(alias) {
				return emoji.Emoji, nil
			}
		}

		for _, tag := range emoji.Tags {
			if re.MatchString(tag) {
				return emoji.Emoji, nil
			}
		}
	}

	return "", errors.New("no matching emoji")
}
