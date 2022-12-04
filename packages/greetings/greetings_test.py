import unittest

from packages.greetings import greetings


class GreetingsTest(unittest.TestCase):
    def test_greeting(self):
        self.assertTrue(greetings.greeting())
