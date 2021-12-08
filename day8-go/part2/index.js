const { readFileSync } = require("fs");

const data = readFileSync('../input', { encoding: 'utf-8' }).trim().split('\n').map(line => {
	const [patterns, fourDigitNum] = line.split('|')
	return { patterns: patterns.trim().split(" "), fourDigitNum: fourDigitNum.trim().split(" ") }
});


class Num {
	constructor(chars, value) {
		this.chars = new Set(chars);
		this.value = value;
	}

	subtract(otherNum) {
		if (this.chars.size < otherNum.chars.size) return;
		for (const iterator of this.chars) {
			if (!otherNum.chars.has(iterator)) {
				return iterator;
			}
		}
		return ''
	}

	subtractSegments(...chars) {
		const remainingChars = new Set(this.chars);
		for (const char of chars) {
			remainingChars.delete(char)
		}
		return Array.from(remainingChars)
	}

	addSegment(...chars) {
		const currentChars = new Set(this.chars);
		for (const char of chars) {
			currentChars.add(char);
		}
		return currentChars.size;
	}

	static add(...nums) {
		const baseNum = new Num('');
		nums.forEach((num) => {
			for (const iterator of num.chars) {
				baseNum.addChar(iterator)
			}
		})
		return baseNum
	}

	addChar(char) {
		this.chars.add(char);
	}

	isNum(chars) {
		const currentChars = new Set(this.chars);
		if (chars.length !== currentChars.size) return false;

		for (const char of chars) {
			currentChars.delete(char);
		}

		return currentChars.size === 0;
	}
}


// const { patterns, fourDigitNum } = {
// 	patterns: ['acedgfb', 'cdfbe', 'gcdfa', 'fbcad', 'dab', 'cefabd', 'cdfgeb', 'eafb', 'cagedb', 'ab'],
// 	fourDigitNum: ['cdfeb', 'fcadb', 'cdfeb', 'cdbaf',]
// };

let sum = 0;

for (const { patterns, fourDigitNum } of data) {
	sum += parseNum(patterns, fourDigitNum)
}

console.log(sum);




function parseNum(patterns, fourDigitNum) {

	const segments = new Array(7);

	let zero;
	let one;
	let two;
	let three;
	let four;
	let five;
	let six;
	let seven;
	let eight;
	let nine;
	let lenSixChars = [];
	let lenFiveChars = [];
	
	patterns.forEach(pattern => {
		switch (pattern.length) {
			case 2:
				one = new Num(pattern, '1');
				break;
			case 4:
				four = new Num(pattern, '4');
				break;
			case 3:
				seven = new Num(pattern, '7');
				break;
			case 7:
				eight = new Num(pattern, '8');
				break;
			case 6:
				lenSixChars.push(pattern);
				break;
			case 5:
				lenFiveChars.push(pattern);
				break;
		}
	})
	
	// filter to get segment 0.
	segments[0]=seven.subtract(one);
	
	// filter segment 6 and get number 9.
	const remainingChars = lenSixChars.filter(chars => {
		const addedVal = Num.add(seven, one, four);
		const str = new Set(chars)
		for (let char of chars) {
			if (addedVal.chars.has(char)) {
				addedVal.chars.delete(char);
				str.delete(char);
			}
		}
		if (addedVal.chars.size === 0 && str.size === 1) {
			nine = new Num(chars, '9')
			segments[6] = Array.from(str)[0];
			return false
		}
		return true;
	})
	
	// segment 4 in.
	segments[4] = eight.subtract(nine)
	
	// filter zero and 6 and set segment to 3.
	remainingChars.forEach(chars => {
		const num = new Num(chars);
		const res = eight.subtract(Num.add(num, one));
		if ( res !== '') {
			zero = new Num(chars, '0');
			segments[3] = res
		} else {
			six = new Num(chars, '6');
		}
	})
	
	segments[2] = eight.subtract(six);
	segments[5] = one.subtractSegments(segments[2])[0];
	segments[1] = four.subtractSegments(segments[2], segments[3], segments[5])[0];
	
	
	lenFiveChars.forEach(chars => {
		const num = new Num(chars);
		if (num.addSegment(segments[2], segments[4]) === 7) {
			// number 5.
			five = new Num(chars, '5');
		} else if (num.addSegment(segments[1], segments[4]) === 7) {
			// number 3
			three = new Num(chars, '3')
		} else {
			// number 2
			two = new Num(chars, '2');
		}
	})
	
	const nums = [zero, one, two, three, four, five, six, seven, eight, nine]
	
	const num = parseInt(fourDigitNum.map(numChars => nums.find(num => num.isNum(numChars)).value).join(""));
	return num
}
