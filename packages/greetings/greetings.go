package greetings

import "math/rand"

type Word struct {
	Locale string
	Word   string
}

var helloEn = Word{
	Locale: "en",
	Word:   "Hello",
}

var hellos = []Word{
	helloEn,
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

func Greeting() string {
	randomIndex := rand.Intn(len(hellos))
	return hellos[randomIndex].Word
}
