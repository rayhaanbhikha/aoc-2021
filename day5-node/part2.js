const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');

class Vector {
  constructor(p1, p2) {
    this.p1 = p1;
    this.p2 = p2;
  }

  get coords() {
    const coords = [];
    if (this.p1.x === this.p2.x) {
      // vertical line
      const min = Math.min(this.p1.y, this.p2.y);
      const max = Math.max(this.p1.y, this.p2.y);
      for (let i = min; i <= max; i++) {
        coords.push({ x: this.p1.x, y: i })
      }
    } else if (this.p1.y === this.p2.y) {
      // horizontal line
      const min = Math.min(this.p1.x, this.p2.x);
      const max = Math.max(this.p1.x, this.p2.x);
      for (let i = min; i <= max; i++) {
        coords.push({ x: i, y: this.p1.y })
      }
    }
    else {
      const m = (this.p2.y - this.p1.y) / (this.p2.x - this.p1.x);
      let newX = this.p1.x;
      let newY = this.p1.y;
      let yToggle = this.p1.y > this.p2.y ? -1 : 1
      let xToggle = this.p1.x > this.p2.x ? -1 : 1
      do {
        coords.push({ x: newX, y: newY })
        newX += (1 * yToggle * m)
        newY += (1 * xToggle * m)
      } while (newX != this.p2.x && newY != this.p2.y);
      coords.push({ x: newX, y: newY })
    }
    return coords
  }
}

let maxN = 0;

const vectors = data.map(val => {
  const [v1, v2] = val.split(" -> ");
  const [x1, y1] = v1.split(",");
  const [x2, y2] = v2.split(",");
  const [nx1, ny1, nx2, ny2] = [Number(x1), Number(y1),  Number(x2), Number(y2)]
  maxN = Math.max(maxN, nx1, ny1, nx2, ny2);
  return new Vector({ x: nx1, y: ny1 }, { x: nx2, y: ny2 })
})

class Grid {
  constructor(size) {
    this.size = size;
    this.grid = [];
    this.populate();
  }

  populate() {
    for (let i = 0; i <= this.size; i++) {
      this.grid.push([])
      for (let j = 0; j <= this.size; j++) {
        this.grid[i].push(0);
      }
    }
  }

  increment(x, y) {
    this.grid[y][x] += 1
  }

  print() {
    for (let i = 0; i <= this.size; i++) {
      console.log("\n")
      for (let j = 0; j <= this.size; j++) {
        process.stdout.write(`${this.grid[i][j]} `)
      }
    }
    console.log("\n")
  }

  valsOver(maxNum) {
    let sum = 0;
    for (let i = 0; i <= this.size; i++) {
      for (let j = 0; j <= this.size; j++) {
        if (this.grid[i][j] >= maxNum ) sum++
      }
    }
    return sum
  }
}

// console.log(new Vector({ x: 1, y: 5 }, { x: 5, y: 1}).coords)
// console.log(new Vector({ x: 9, y: 7 }, { x: 7, y: 9}).coords)
// console.log(new Vector({ x: 5, y: 3}, { x: 7, y: 5 }).coords)
// console.log(new Vector({ x: 5, y: 5}, { x: 8, y: 2 }).coords)
// console.log(new Vector({ x: 8, y: 2 }, { x: 5, y: 5}).coords)
// console.log(new Vector({ x: 7, y: 9 }, { x: 9, y: 7}).coords)
// console.log(new Vector({ x: 0, y: 1 }, { x: 3, y: 3}).coords)

const grid = new Grid(maxN);

vectors.forEach(vector => {
  vector.coords.forEach(({ x, y }) => {
    grid.increment(x, y)
  })
})

// grid.print();

console.log(grid.valsOver(2))
