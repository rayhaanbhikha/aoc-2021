const { readFileSync } = require("fs");

const lines = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n').map(l => l.split(""));

const closingBracket = {
  "(": ")",
  "{": "}",
  "[": "]",
  "<": ">"
}

const bracketScore = {
  ")": 1,
  "]": 2,
  "}": 3,
  ">": 4,
}

const closingBrackets = new Set([")", "}", "]", ">"]);

function filterIncompleteLines(line) {
  const bracketStack = [];

  for (const bracket of line) {
    const lastBracket = bracketStack.at(-1);
    const matchingClosingBracket = closingBracket[lastBracket];
    
    if (closingBrackets.has(bracket)) {
      if (bracket !== matchingClosingBracket) {
        // it's a corrupted line bracket
        return null;
      }
      bracketStack.pop();
      continue;
    }

    bracketStack.push(bracket);
  }

  if (bracketStack.length > 0) {
    return bracketStack;
  }

  return null;
}

function computeScore(line) {
  const filteredLine = filterIncompleteLines(line)
  if (!filteredLine) return -1;
  return filteredLine.reverse().reduce((sum, openBracket) => {
    const matchingClosingBracket = closingBracket[openBracket];
    sum *= 5
    sum += bracketScore[matchingClosingBracket]
    return sum
  }, 0)
}

const scores = lines.map(computeScore).filter(score => score >= 0).sort((a, b) => a - b);

const middleIndex = Math.floor(scores.length / 2);

console.log(scores[middleIndex])

