import { cleanInput } from "./repl";
import { describe, expect, test } from "vitest";

describe.each([
  {
    input: "   hello   world   ",
    expected: ["hello", "world"],
  },
  {
    input: "Hi how are you",
    expected: ["hi", "how", "are", "you"],
  },
  {
    input: "AksHuaLly, it's GNU Linux",
    expected: ["akshually,", "it's", "gnu", "linux"],
  },
])("cleanInput($input)", ({ input, expected }) => {
  test(`Expected: ${expected}`, () => {
    let actual = cleanInput(input);

    expect(actual).toHaveLength(expected.length);
    for (const i in expected) {
      expect(actual[i]).toBe(expected[i]);
    }
  });
});
