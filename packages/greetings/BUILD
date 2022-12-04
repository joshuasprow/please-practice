python_library(
    name="greetings",
    srcs=["greetings.py"],
    visibility=["//apps/...", "//packages/..."],
    deps=["//pls-deps/python:python-i18n"],
)

python_test(
    name="greetings_test",
    srcs=["greetings_test.py"],
    # Here we have used the shorthand `:greetings` label format. This format can be used to refer to a rule in the same
    # package and is shorthand for `//packages/greetings:greetings`.
    deps=[":greetings"],
)
