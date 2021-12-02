const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');

class Submarine {
  constructor() {
    this.x = 0
    this.y = 0
  }

  move(direction, units) {
    switch (direction) {
      case 'forward':
        this.x += units
        break;
      case 'down':
        this.y += units;
        break;
      case 'up':
        this.y -= units;
        break;
    }
  }

  result() {
    return this.x * this.y
  }
}

const submarine = new Submarine();

for (const item of data) {
  const [direction, units] = item.split(' ');
  submarine.move(direction, Number(units));
}

console.log(submarine.result());
