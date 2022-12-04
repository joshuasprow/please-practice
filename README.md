### `/root/`

`.gitignore`: ignore build files, secrets, temporary files

`.plzconfig`: configuration for [Please](https://please.build/) build tool

`pleasew`: binary that will download and install the correct version of Please automagically

`poetry.lock, pyproject.toml`: we're using [Poetry](https://python-poetry.org/) only to set the project's Python version and maintain virtual environments, run dev tools, etc.

[`BUILD`](https://please.build/language.html)  files: configs in directories that are run/built with Please