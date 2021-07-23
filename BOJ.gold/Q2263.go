package main

import (
	"fmt"
)

type node struct {
	val       int
	leftNode  *node
	rightNode *node
}

var n int
var inOrder = []int{}
var postOrder = []int{}
var rootNode *node

func printNode(node *node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.val)

	printNode(node.leftNode)
	printNode(node.rightNode)
}

func split(parent *node, tree []int, standard int) {
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
	n--
	nextStandard := postOrder[n-1]

	for _, val := range rightTree {
		if val == nextStandard {
			nowNode := node{val: nextStandard}
			parent.rightNode = &nowNode
			split(&nowNode, rightTree, nextStandard)

		}
	}
	for _, val := range leftTree {
		if val == nextStandard {
			nowNode := node{val: nextStandard}
			parent.leftNode = &nowNode
			split(&nowNode, leftTree, nextStandard)
		}
	}

}

func main() {
	//fmt.Scan(&n)
	// inOrder = make([]int, 12)
	// postOrder = make([]int, 12)

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)
	// for i := 0; i < n; i++ {
	// 	scanner.Scan()
	// 	tmp, _ := strconv.Atoi(scanner.Text())
	// 	inOrder[i] = tmp
	// }
	// for i := 0; i < n; i++ {
	// 	scanner.Scan()
	// 	tmp, _ := strconv.Atoi(scanner.Text())
	// 	postOrder[i] = tmp
	// }
	n = 12
	inOrder = append(inOrder, 11)
	inOrder = append(inOrder, 8)
	inOrder = append(inOrder, 4)
	inOrder = append(inOrder, 9)
	inOrder = append(inOrder, 12)
	inOrder = append(inOrder, 2)
	inOrder = append(inOrder, 5)
	inOrder = append(inOrder, 1)
	inOrder = append(inOrder, 6)
	inOrder = append(inOrder, 3)
	inOrder = append(inOrder, 7)
	inOrder = append(inOrder, 10)
	postOrder = append(postOrder, 11)
	postOrder = append(postOrder, 8)
	postOrder = append(postOrder, 12)
	postOrder = append(postOrder, 9)
	postOrder = append(postOrder, 4)
	postOrder = append(postOrder, 5)
	postOrder = append(postOrder, 2)
	postOrder = append(postOrder, 6)
	postOrder = append(postOrder, 10)
	postOrder = append(postOrder, 7)
	postOrder = append(postOrder, 3)
	postOrder = append(postOrder, 1)
	split(nil, inOrder, postOrder[n-1])
	printNode(rootNode)
}
