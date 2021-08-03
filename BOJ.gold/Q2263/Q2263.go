package q2263

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	val       int
	leftNode  *node
	rightNode *node
}

var stdout *bufio.Writer
var n int
var inOrder = []int{}
var postOrder = []int{}
var rootNode *node

func printNode(node *node) {
	if node == nil {
		return
	}
	stdout.WriteString(strconv.Itoa(node.val) + " ")
	printNode(node.leftNode)
	printNode(node.rightNode)
}

func split(parent *node, tree []int, standard int, standardTree []int) {
	var pivot int
	if parent == nil {
		rootNode = &node{val: standard}
		parent = rootNode
	}
	if len(tree) == 0 {
		return
	}
	for i, val := range tree {
		if val == standard {
			pivot = i
			break
		}
	}
	leftTree := tree[:pivot]
	rightTree := tree[pivot+1:]

	rightStandardTree := standardTree[len(standardTree)-len(rightTree)-1 : len(standardTree)-1]
	leftStandardTree := standardTree[:len(standardTree)-len(rightTree)-1]
	for _, val := range rightTree {
		nextStandard := rightStandardTree[len(rightStandardTree)-1]
		if val == nextStandard {
			nowNode := node{val: nextStandard}
			parent.rightNode = &nowNode
			split(&nowNode, rightTree, nextStandard, rightStandardTree)

		}
	}
	for _, val := range leftTree {
		nextStandard := leftStandardTree[len(leftStandardTree)-1]

		if val == nextStandard {
			nowNode := node{val: nextStandard}
			parent.leftNode = &nowNode
			split(&nowNode, leftTree, nextStandard, leftStandardTree)
		}
	}

}

func Q2263() {
	fmt.Scan(&n)
	inOrder = make([]int, n)
	postOrder = make([]int, n)
	stdout = bufio.NewWriter(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp, _ := strconv.Atoi(scanner.Text())
		inOrder[i] = tmp
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp, _ := strconv.Atoi(scanner.Text())
		postOrder[i] = tmp
	}

	split(nil, inOrder, postOrder[n-1], postOrder)
	printNode(rootNode)
	stdout.Flush()
}
