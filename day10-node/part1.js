const { readFileSync } = require("fs");

const lines = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n').map(l => l.split(""));

const closingBracket = {
  "(": ")",
  "{": "}",
  "[": "]",
  "<": ">"
}

const bracketScore = {
  ")": 3,
  "]": 57,
  "}": 1197,
  ">": 25137,
}

const closingBrackets = new Set([")", "}", "]", ">"]);

function findCorruptedClosingCharacter(line) {
  const bracketStack = [];

  for (const bracket of line) {
    const lastBracket = bracketStack.at(-1);
    const matchingClosingBracket = closingBracket[lastBracket];

    if (closingBrackets.has(bracket)) {
      if (bracket !== matchingClosingBracket) {
        // corrupted.
        return bracket;
      }
      bracketStack.pop();
      continue;
    }

    bracketStack.push(bracket);
  }

  return null;
}

const mappings = lines.reduce((acc, line) => {
  const res = findCorruptedClosingCharacter(line);
  if (res) {
    acc[res] = acc[res] ? acc[res] + 1 : 1
  }
  return acc
}, {})

const score = Object.entries(mappings).reduce((sum, [bracket, occurrence]) => sum += (bracketScore[bracket] * occurrence), 0)

console.log(score);
