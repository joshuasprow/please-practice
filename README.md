| <h2>`/`</h2>                                               |                                                                                                                                                    |
| ---------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `.gitignore`                                               | ignore build files, secrets, temporary files                                                                                                       |
| `poetry.lock`<br>`pyproject.toml`                          | we're using [`poetry`](https://python-poetry.org/) only to set the project's Python version and maintain virtual environments, run dev tools, etc. |
| `pleasew`                                                  | binary that will download and install the correct version of [`please`](https://please.build) automagically                                        |
| `**/`[`BUILD`](https://please.build/language.html)` files` | configs in directories that are run/built with `please`                                                                                            |

| <h2>`/plz/`</h2> |                                                                                              |
| ---------------- | -------------------------------------------------------------------------------------------- |
| `python/`        | `BUILD` files that configure third party dependencies, all of which are visible project-wide |