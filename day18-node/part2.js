const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');

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
    }

    magnitude() {
        if (this.leafNode) { // single value.
            return this.val;
        }

        // node just before pair.
        if (this.left?.leafNode && this.right?.leafNode) {
            return (3 * this.left.val) + (2 * this.right.val);
        }

        return 3 * this.left.magnitude() + 2 * this.right.magnitude();
    }
}

let result = 0;

for (const pairString1 of data) {
    for (const pairString2 of data) {
        const n1 = Node.fromString(pairString1);
        const n2 = Node.fromString(pairString2);
        const res = Node.addNodes(n1, n2);
        res.reduceNodes();
        result = Math.max(res.magnitude(), result);        
    }
}

console.log(result)