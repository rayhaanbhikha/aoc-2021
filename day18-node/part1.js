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
    constructor(parent, pairsString, level) {
        this.parent = parent;
        this.level = level;
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
        return new Node(null, pairsString, 0)
    }

    static addNodes(node1, node2) {
        const parentNode = new Node(null, "root", 0);
        parentNode.left = node1;
        node1.parent = parentNode;
        node1.incrementLevel();
        parentNode.right = node2;
        node2.parent = parentNode;
        node2.incrementLevel();
        return parentNode;
    }

    incrementLevel() {
        this.level++
        this.left?.incrementLevel();
        this.right?.incrementLevel();
    }

    parse(data) {
        const { left, right } = parsePair(data)
        if (left) {
            this.left = new Node(this, left, this.level + 1);
        }

        if (right) {
            this.right = new Node(this, right, this.level + 1);
        }
    }

    split() {
        const answer = this.val / 2
        this.leafNode = false;
        this.val = 0;
        this.left = new Node(this, Math.floor(answer), this.level+1);
        this.right = new Node(this, Math.ceil(answer), this.level+1);
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
    }

    printf(res) {
        if (this.leafNode) {
            res.push({ level: this.level, val: this.val, leafNode: true });
            return
        }
        res.push({ level: this.level, branch: 'L', val: 'x', leafNode: false })
        this.left?.printf(res);
        res.push({ level: this.level, branch: 'R', val: 'x', leafNode: false })
        this.right?.printf(res);
    }

    treeVals() {
        const res = [];
        this.printf(res);
        return res.filter(v => v.leafNode).map(v => v.val).join(",");
    }

    print() {
        const res = [];
        this.printf(res);
        console.log(res.sort((a, b) => a.level - b.level));
        return ;
    }

    explodePairs() {
        if(this.left?.explodePairs()) {
            return true;
        }

        if (this.level >= 4 && this.left?.leafNode && this.right?.leafNode) {
            // just above leaf nodes.
            this.explode();
            return true;
        }
     
        if (this.right?.explodePairs()) {
            return true
        }

        return false;
    }

    splitNum() {
        if(this.left?.splitNum()) {
            return true;
        }

        if (this.leafNode && this.val >= 10) {
            this.split();
            return true
        }
     
        if (this.right?.splitNum()) {
            return true
        }

        return false;
    }

    explodePairsRepeat() {
        const res = this.explodePairs();
        if (res) {
            return this.explodePairsRepeat();
        }
    }

    splitRepeat() {
        const res = this.splitNum();
        if (res) {
            return this.splitRepeat();
        }
    }

    reduceNodes() {
        const res = this.explodePairsRepeat() || this.splitNum();
        if (res) {
            return this.reduceNodes();
        }
        // while (this.explodePairsRepeat() || this.splitRepeat()) {}
    }
}

// // ----------------- Start Test Reducer -------------
// const pairs = [
//     ["[[[[[9,8],1],2],3],4]", "0,9,2,3,4"],
//     ["[7,[6,[5,[4,[3,2]]]]]", "7,6,5,7,0"],
//     ["[[6,[5,[4,[3,2]]]],1]", "6,5,7,0,3"],
//     ["[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "3,2,8,0,9,5,7,0"]
// ]

// pairs.forEach((val) => {
//     const [pairsString, expectedResult] = val;
//     const n = Node.fromString(pairsString);
//     n.reduceNodes();
//     assert(n.treeVals() === expectedResult, `reducer failed. String: ${pairsString}, Expected Result${expectedResult} `);
// })
// // ----------------- End Test Reducer -------------

// const n1 = Node.fromString("[[[[4,3],4],4],[7,[[8,4],9]]]")
// const n2 = Node.fromString("[1,1]")
// const sum = Node.addNodes(n1, n2);
// sum.reduceNodes();
// assert(sum.treeVals() === "0,7,4,7,8,6,0,8,1");


// const node1 = Node.fromString("[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]");
// const node2 = Node.fromString("[2,9]");
const node1 = Node.fromString(data[0])
const node2 = Node.fromString(data[1])

// node1.print();
// console.log(node1.treeVals());
// console.log(node2.treeVals());

const res = Node.addNodes(node1, node2);
res.reduceNodes();
console.log(res.treeVals());








// const pairs = ['[1,1]', '[2,2]', '[3,3]', '[4,4]', '[5,5]', '[6,6]'];

// const res = pairs
//     .map(n => Node.fromString(n))
//     .reduce((acc, node) => {
//         return Node.addNodes(acc, node);
//     });

// const n2 = Node.fromString("[[6,[5,[4,[3,2]]]],1]")
// n2.reduceNodes();
// n2.treeVals();

// const t = Node.fromString("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
// t.reduceNodes();
// t.treeVals();

// n1.reduceNodes()
// n1.treeVals()
// console.log("\n\n----\n\n")
// n2.treeVals()
// const res = Node.addNodes(n1, n2);
// res.reduceNodes()
// res.treeVals()
// const result= data
//     .map(pairString => Node.fromString(pairString))
//     .reduce((acc, node) => {
//         console.log(acc, node);
//         const res = Node.addNodes(acc, node)
//         res.reduce();
//         return res;
//     });

