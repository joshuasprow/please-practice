import unittest

from packages.greetings import greetings

# TODO: fix filepath for tests


class GreetingsTest(unittest.TestCase):
    def test_greeting(self):
        self.assertTrue(greetings.greeting())
