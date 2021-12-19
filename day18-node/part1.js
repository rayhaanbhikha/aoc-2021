const { readFileSync } = require("fs");

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
        if (num) {
            this.val = num;
            this.leafNode = true;
        } else {
            this.rawString = pairsString;
            this.parse(pairsString);
            this.leafNode = false;
            this.val = 0;
        }
    }

    static fromString(pairsString) {
        return new Node(null, pairsString, 0)
    }

    static addNodes(node1, node2) {
        const parentNode = new Node(null, "");
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

        if(this.right.leafNode) {
            this.right.updateVal(val);
            return;
        }
        this.right.insertRightMostLeafNode(val);
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

        if(this.left.leafNode) {
            this.left.updateVal(val);
            return;
        }
        this.left.insertLeftMostLeafNode(val);
    }

    updateVal(val) { 
        this.val += val;
        if (this.val >= 10) this.split();
    }

    print() {
        if (this.leafNode) {
            process.stdout.write(`${this.val} `);
            return
        }
        this.left?.print();
        process.stdout.write(",");
        this.right?.print();
    }

    reduce(level=0) {
        if (this.leafNode && this.val >= 10) {
            this.split();
            return
        }

        this.left?.reduce(level+1)

        if (level >= 4 && this.left?.leafNode && this.right?.leafNode) {
            // just above leaf nodes.
            this.explode();
            return;
        }

        this.right?.reduce(level+1)
    }
}

const n1 = Node.fromString("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
const n2 = Node.fromString(data[1])
// console.log(n1.rawString)
n1.print()
// console.log(n1 === n1);
// console.log(n1 === n2);
// console.log("----")
// const res = Node.addNodes(n1, n2);
// console.log(res);
// const res = Node.fromString("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
// const res = Node.fromString("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
// res.reduce()
// res.print();
// const result= data
//     .map(pairString => Node.fromString(pairString))
//     .reduce((acc, node) => {
//         console.log(acc, node);
//         const res = Node.addNodes(acc, node)
//         res.reduce();
//         return res;
//     });

// result.print()


// const node = new Node(null, "[3,[1,5]]");

// console.log(node);
// node.right.explode();
// console.log(node);
// const n = Node.fromString("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
// const n = Node.fromString("[[6,[5,[4,[3,2]]]],1]")
// n.reduce();

// console.log(n);
// n.print()
// console.log(JSON.stringify(n, null, 2))




// const node1 = new Node(null, data);
// const node2 = new Node(null, "[1,1]");
// // node1.print()
// // const node3 = Node.addNodes(node1, node2)

// // console.log(node3);
// const node = new Node(null, "[[15,3],4]")
// node.print()
// // console.log(node)
// // console.log(node.left.split());
// // console.log(node)