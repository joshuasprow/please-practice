import json
import random
from os import path
from typing import NamedTuple

import i18n


class Word(NamedTuple):
    locale: str
    word: str


def read_hellos_from_file() -> list[Word]:
    filepath = path.join("data", "greetings.json")
    objs: list[Word] = []

    with open(filepath, "r") as f:
        objs = json.load(f)

    return [Word(**obj) for obj in objs]


def load_bundle(hellos: list[Word], hello_en: Word):
    for hello in hellos:
        i18n.add_translation(
            hello_en.word,
            hello.word,
            locale=hello.locale,
        )


def greeting():
    hellos = read_hellos_from_file()
    hello_en = next(hello for hello in hellos if hello.locale == "en")

    if hello_en is None:
        raise Exception("No English hello found")

    load_bundle(hellos, hello_en)

    locale = random.choice(hellos).locale

    return i18n.t(hello_en.word, locale=locale)
