import i18n from "i18next";

class Word {
  constructor(locale, word) {
    this.locale = locale;
    this.word = word;
  }
}

const helloEn = new Word("en", "Hello");

const hellos = [
  helloEn,
  new Word("es", "Hola"),
  new Word("fr", "Bonjour"),
  new Word("de", "Hallo"),
  new Word("it", "Ciao"),
  new Word("pt", "Olá"),
  new Word("ru", "Привет"),
  new Word("ja", "こんにちは"),
  new Word("zh", "你好"),
  new Word("ko", "안녕하세요"),
  new Word("ar", "مرحبا"),
  new Word("hi", "नमस्ते"),
  new Word("bn", "হ্যালো"),
  new Word("th", "สวัสดี"),
  new Word("vi", "Xin chào"),
  new Word("id", "Halo"),
  new Word("ms", "Hai"),
];

/** @type {import("i18next").Resource} */
const resources = {};

for (const hello of hellos) {
  resources[hello.locale] = { [helloEn.word]: hello.word };
}

i18n.init({ resources });

/** @returns {string} */
export const greeting = () => {
  const randomHello = hellos[Math.floor(Math.random() * hellos.length)];
  const randomLanguage = randomHello.locale;
  const hello = i18n.getDataByLanguage(randomLanguage)[helloEn.word];

  return hello;
};

console.log(greeting());
