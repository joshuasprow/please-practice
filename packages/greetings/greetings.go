package greetings

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const messageID = "hello"

type Word struct {
	Locale string `json:"locale"`
	Word   string `json:"word"`
}

var bundle *i18n.Bundle

func readHellosFromFile() ([]Word, error) {
	path := filepath.Join("..", "..", "data", "greetings.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hellos := []Word{}

	if err := json.Unmarshal(data, &hellos); err != nil {
		return nil, err
	}

	return hellos, nil
}

func loadBundle(hellos []Word) (*i18n.Bundle, error) {
	b := i18n.NewBundle(language.English)

	for _, hello := range hellos {
		msg := &i18n.Message{
			ID:    messageID,
			Other: hello.Word,
		}

		tag, err := language.Parse(hello.Locale)
		if err != nil {
			return nil, err
		}

		if err := b.AddMessages(tag, msg); err != nil {
			return nil, err
		}
	}

	return b, nil
}

func getRandomHello(hellos []Word) (string, error) {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(hellos))
	locale := hellos[index].Locale

	return i18n.
		NewLocalizer(bundle, locale).
		Localize(&i18n.LocalizeConfig{MessageID: messageID})
}

func Greeting() (string, error) {
	var err error

	hellos, err := readHellosFromFile()
	if err != nil {
		return "", err
	}

	if bundle == nil {
		bundle, err = loadBundle(hellos)
		if err != nil {
			return "", err
		}
	}

	return getRandomHello(hellos)
}
