const { readFileSync } = require("fs");
const assert = require('assert');

const data = readFileSync('./sample', { encoding: 'utf-8' }).trim().split('\n');
// const data = "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"

function parsePair(data) {
    const insideData = data.slice(1, data.length - 1)
    let brackets = 0
    for (let i=0; i<insideData.length; i++) {
        const char = insideData[i]
        if (brackets == 0 && char == ",") {
            return { left: insideData.slice(0, i), right: insideData.slice(i+1)}
        }
        if (char == "[") {
            brackets++
            continue
        }

        if (char == "]") {
            brackets--
        }
    }
    return { left: '', right: '' }
}

class Node {
    constructor(parent, pairsString) {
        this.parent = parent;
        const num = Number(pairsString);

        if (pairsString == "root") {
            this.val = 0;
            this.leafNode = false;
        } else if (num !== 0 && (isNaN(num) || num == "")) {
            this.rawString = pairsString;
            this.parse(pairsString);
            this.leafNode = false;
            this.val = 0;
        } else {
            this.val = num;
            this.leafNode = true;
        }
    }

    static fromString(pairsString) {
        return new Node(null, pairsString)
    }

    static addNodes(node1, node2) {
        const parentNode = new Node(null, "root");
        parentNode.left = node1;
        node1.parent = parentNode;
        parentNode.right = node2;
        node2.parent = parentNode;
        return parentNode;
    }

    parse(data) {
        const { left, right } = parsePair(data)
        if (left) {
            this.left = new Node(this, left);
        }

        if (right) {
            this.right = new Node(this, right)
        }
    }

    split() {
        const answer = this.val / 2
        this.leafNode = false;
        this.val = 0;
        this.left = new Node(this, Math.floor(answer));
        this.right = new Node(this, Math.ceil(answer));
    }

    explode() {
        this.findLeftMostInsertionNode(this.left.val);
        this.findRightMostInsertionNode(this.right.val);
        this.leafNode = true;
        this.val = 0;
        this.left = null;
        this.right = null;
    }

    findLeftMostInsertionNode(val) {
        if (this.parent == null) return;

        if (this.parent.left == this) {
            this.parent.findLeftMostInsertionNode(val);
            return;
        }

        this.parent.left.insertRightMostLeafNode(val);
        return
    }

    insertRightMostLeafNode(val) {
        if (this.leafNode) {
            this.updateVal(val);
            return;
        }

        if(this.right?.leafNode) {
            this.right.updateVal(val);
            return;
        }
        this.right?.insertRightMostLeafNode(val);
    }

    findRightMostInsertionNode(val) {
        if (this.parent == null) return;

        if (this.parent.right == this) {
            this.parent.findRightMostInsertionNode(val);
            return;
        }

        // it's flipped.
        this.parent.right.insertLeftMostLeafNode(val);
        return;
    }

    insertLeftMostLeafNode(val) {
        if (this.leafNode) {
            this.updateVal(val);
            return;
        }

        if(this.left?.leafNode) {
            this.left.updateVal(val);
            return;
        }
        this.left?.insertLeftMostLeafNode(val);
    }

    updateVal(val) { 
        this.val += val;
        if (this.val >= 10) this.split();
    }

    printf(res) {
        if (this.leafNode) {
            res.push(this.val);
            // process.stdout.write(`${this.val}`);
            return
        }
        this.left?.printf(res);
        // process.stdout.write(",");
        this.right?.printf(res);
    }

    print() {
        const res = [];
        this.printf(res);
        return res.join(",");
    }

    reduce(level=0) {
        if (this.leafNode && this.val >= 10) {
            this.split();
            return true
        }

        if(this.left?.reduce(level+1)) {
            return true;
        }

        if (level >= 4 && this.left?.leafNode && this.right?.leafNode) {
            // just above leaf nodes.
            this.explode();
            return true;
        }

        if (this.right?.reduce(level+1)) {
            return true;
        }

        return false;
    }

    reduceNodes() {
        let isReducing = true;
        do {
            isReducing = this.reduce();
        } while (isReducing)
    }
}

// ----------------- Start Test Reducer -------------
const pairs = [
    ["[[[[[9,8],1],2],3],4]", "0,9,2,3,4"],
    ["[7,[6,[5,[4,[3,2]]]]]", "7,6,5,7,0"],
    ["[[6,[5,[4,[3,2]]]],1]", "6,5,7,0,3"],
    ["[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "3,2,8,0,9,5,7,0"]
]

pairs.forEach((val) => {
    const [pairsString, expectedResult] = val;
    const n = Node.fromString(pairsString);
    n.reduceNodes();
    assert(n.print() === expectedResult, `reducer failed. String: ${pairsString}, Expected Result${expectedResult} `);
})
// ----------------- End Test Reducer -------------


// const pairs = ['[1,1]', '[2,2]', '[3,3]', '[4,4]', '[5,5]', '[6,6]'];

// const res = pairs
//     .map(n => Node.fromString(n))
//     .reduce((acc, node) => {
//         return Node.addNodes(acc, node);
//     });

// res.reduceNodes()
// res.print()
// console.log(res.left)
// console.log(res.right)

// const n2 = Node.fromString("[[6,[5,[4,[3,2]]]],1]")
// n2.reduceNodes();
// n2.print();

// const t = Node.fromString("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
// t.reduceNodes();
// t.print();

// const n1 = Node.fromString(data[0])
// const n2 = Node.fromString(data[1])
// n1.reduceNodes()
// n1.print()
// console.log("\n\n----\n\n")
// n2.print()
// const res = Node.addNodes(n1, n2);
// res.reduceNodes()
// res.print()
// const result= data
//     .map(pairString => Node.fromString(pairString))
//     .reduce((acc, node) => {
//         console.log(acc, node);
//         const res = Node.addNodes(acc, node)
//         res.reduce();
//         return res;
//     });

