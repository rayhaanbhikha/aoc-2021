const fs = require('fs');
const data = fs.readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');



const part1 = () => {
  let numIncreased = 0;

  for (let i = 1; i < data.length; i++) {
    const lastNum = Number(data[i - 1]);
    if (Number(data[i]) > lastNum) {
      numIncreased++
    }
  }
  
  console.log(numIncreased);
}

const part2 = () => {
  let numIncreased = 0;

  for (let i = 1; i < data.length - 2; i++) {
    const lastNum = Number(data[i - 1]) + Number(data[i]) + Number(data[i + 1]);
    const currentNum = Number(data[i]) + Number(data[i + 1]) + Number(data[i + 2]);
    if (currentNum > lastNum) {
      numIncreased++
    }
  }
  
  console.log(numIncreased);
}

part1();
part2();
