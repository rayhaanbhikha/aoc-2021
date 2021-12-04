const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');

class BingoBoard {
  constructor() {
    this.grid = [];
  }

  markNum(num) {
    for (let i = 0; i < 5; i++) {
      for (let j = 0; j < 5; j++) {
        if (num === (this.grid[i][j]).num) {
          this.grid[i][j].marked = true
        }
      }
    }
  }

  addRow(numbers) {
    this.grid.push(numbers.split(" ").filter(n => n !== "").map(n => ({ num: Number(n), marked: false })))
  }

  hasWon() {
    // check rows
    for (let rowIndex = 0; rowIndex < 5; rowIndex++) {
      if (this.checkRowComplete(rowIndex)) return true;
    }

    // check columns
    for (let colIndex = 0; colIndex < 5; colIndex++) {
      if (this.checkColumnComplete(colIndex)) return true;
    }

    return false
  }

  checkRowComplete(rowIndex) {
    for (const nums of this.grid[rowIndex]) {
      if (!nums.marked) return false;
    }
    return true
  }

  checkColumnComplete(index) {
    for (let i = 0; i < 5; i++) {
      if (!this.grid[i][index].marked) return false;
    }
    return true
  }

  sumOfUnmarkedNums() {
    let sum = 0;
    for (const row of this.grid) {
      for (const num of row) {
        if (!num.marked) {
          sum += num.num
        }
      }
    }
    return sum;
  }
}

const bingoBoards = [];

const bingoNumbers = data.shift().split(",").map(Number)

for (const bingoDataRow of data) {
  if (bingoDataRow == "") {
    bingoBoards.push(new BingoBoard());
  } else {
    bingoBoards.at(-1).addRow(bingoDataRow);
  }
}

for (const bingoNum of bingoNumbers) {
  bingoBoards.forEach(bingoBoard => bingoBoard.markNum(bingoNum))
  for (const bingoBoard of bingoBoards) {
    if (bingoBoard.hasWon()) {
      console.log(bingoBoard.sumOfUnmarkedNums() * bingoNum)
      return
    }
  }
}
