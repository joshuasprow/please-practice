import i18n
import random
from typing import NamedTuple


class Word(NamedTuple):
    locale: str
    word: str


hello_en = Word("en", "Hello")

hellos = [
    hello_en,
    Word("es", "Hola"),
    Word("fr", "Bonjour"),
    Word("de", "Hallo"),
    Word("it", "Ciao"),
    Word("pt", "Olá"),
    Word("ru", "Привет"),
    Word("ja", "こんにちは"),
    Word("zh", "你好"),
    Word("ko", "안녕하세요"),
    Word("ar", "مرحبا"),
    Word("hi", "नमस्ते"),
    Word("bn", "হ্যালো"),
    Word("th", "สวัสดี"),
    Word("vi", "Xin chào"),
    Word("id", "Halo"),
    Word("ms", "Hai"),
]

for hello in hellos:
    i18n.add_translation(
        hello_en.word,
        hello.word,
        locale=hello.locale,
    )


def greeting():
    locale = random.choice(hellos).locale
    hello = i18n.t(hello_en.word, locale=locale)

    return hello
