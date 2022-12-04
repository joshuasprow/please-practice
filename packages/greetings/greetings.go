package greetings

import (
	"math/rand"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const messageID = "hello"

var hellos = []struct {
	Locale string
	Word   string
}{
	{Locale: "en", Word: "Hello"},
	{Locale: "es", Word: "Hola"},
	{Locale: "fr", Word: "Bonjour"},
	{Locale: "de", Word: "Hallo"},
	{Locale: "it", Word: "Ciao"},
	{Locale: "pt", Word: "Olá"},
	{Locale: "ru", Word: "Привет"},
	{Locale: "ja", Word: "こんにちは"},
	{Locale: "zh", Word: "你好"},
	{Locale: "ko", Word: "안녕하세요"},
	{Locale: "ar", Word: "مرحبا"},
	{Locale: "hi", Word: "नमस्ते"},
	{Locale: "bn", Word: "হ্যালো"},
	{Locale: "th", Word: "สวัสดี"},
	{Locale: "vi", Word: "Xin chào"},
	{Locale: "id", Word: "Halo"},
	{Locale: "ms", Word: "Hai"},
}

var bundle *i18n.Bundle

func loadBundle() (*i18n.Bundle, error) {
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

func getRandomHello() (string, error) {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(hellos))
	locale := hellos[index].Locale

	return i18n.
		NewLocalizer(bundle, locale).
		Localize(&i18n.LocalizeConfig{MessageID: messageID})
}

func Greeting() (string, error) {
	var err error

	if bundle == nil {
		bundle, err = loadBundle()
		if err != nil {
			return "", err
		}
	}

	return getRandomHello()
}
