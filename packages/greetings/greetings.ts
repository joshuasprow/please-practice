import { readFileSync } from "fs";
import i18n, { Resource } from "i18next";
import { join } from "path";

class Word {
  locale: string;
  word: string;

  constructor(locale: string, word: string) {
    this.locale = locale;
    this.word = word;
  }
}

const readHellosFromFile = () => {
  const path = join(__dirname, "..", "..", "data", "greetings.json");
  const data = readFileSync(path, "utf8");

  const hellos: Word[] = [];

  for (const obj of JSON.parse(data)) {
    const hello = new Word(obj.locale, obj.word);

    hellos.push(hello);
  }

  return hellos;
};

const loadBundle = (hellos: Word[], helloEn: Word) => {
  const resources: Resource = {};

  for (const hello of hellos) {
    resources[hello.locale] = { [helloEn.word]: hello.word };
  }

  i18n.init({ resources });
};

export const greeting = () => {
  const hellos = readHellosFromFile();
  const helloEn = hellos.find((hello) => hello.locale === "en");

  if (!helloEn) {
    throw new Error("English greeting not found");
  }

  loadBundle(hellos, helloEn);

  const randomHello = hellos[Math.floor(Math.random() * hellos.length)];
  const randomLanguage = randomHello.locale;

  return i18n.getDataByLanguage(randomLanguage)[helloEn.word];
};
