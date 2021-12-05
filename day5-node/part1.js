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
    return coords
  }
}

let maxN = 0;

const vectors = data.map(val => {
  const [v1, v2] = val.split(" -> ");
  const [x1, y1] = v1.split(",");
  const [x2, y2] = v2.split(",");
  const [nx1, ny1, nx2, ny2] = [Number(x1), Number(y1),  Number(x2), Number(y2)]
  const currentMaxX = Math.max(nx1, nx2);
  const currentMaxY = Math.max(ny1, ny2);
  const currentMax = Math.max(currentMaxX, currentMaxY)
  maxN = Math.max(maxN, currentMax);
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



const grid = new Grid(maxN);

vectors.forEach(vector => {
  vector.coords.forEach(({ x, y }) => {
    grid.increment(x, y)
  })
})

// console.log(grid.print());

console.log(grid.valsOver(2))
