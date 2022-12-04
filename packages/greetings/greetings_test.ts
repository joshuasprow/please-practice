import * as assert from "node:assert/strict";
import test from "node:test";
import { greeting } from "./greetings.js";

test("greeting returns a string", () => {
  assert.strictEqual(typeof greeting(), "string");
});
