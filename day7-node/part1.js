const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split(',').map(Number);

class Crabs {
  constructor(input) {
    this.crabs = {};
    for (const crabPosition of input) {
      const copies = this.crabs[crabPosition];
      this.crabs[crabPosition] = copies ? copies + 1 : 1
    }
  }

  computeFuelCost(num) {
    let fuelCost = 0;
    for (const crab in this.crabs) {
      const copies = this.crabs[crab];
      const diff = Math.abs(num - Number(crab));
      fuelCost += (diff * copies);
    }
    return fuelCost;
  }
}

const crabPositions = Array.from(new Set(data.sort((a, b) => a - b)));
const crabs = new Crabs(data);

const minimumFuel = (crabPositions, crabs) => {
  let leftIndex = 0;
  let rightIndex = crabPositions.length - 1;
  while (rightIndex > leftIndex) {
    const fuelCostLeft = crabs.computeFuelCost(leftIndex);
    const fuelCostRight = crabs.computeFuelCost(rightIndex);
    if (fuelCostRight > fuelCostLeft) {
      rightIndex--
    } else {
      leftIndex++
    }
  }

  return crabs.computeFuelCost(leftIndex);
}

console.log(minimumFuel(crabPositions, crabs))
