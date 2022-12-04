import os
from os import path

from packages.greetings import greetings


def main():
    greetings_file = os.getenv(
        "GREETINGS_PATH",
        path.join("data", "greetings.json"),
    )

    hello = greetings.greeting(greetings_file)

    print(f"{hello} World!")


if __name__ == "__main__":
    main()
