package q10868

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var segTree []int
var arr []int

type query struct {
	start int
	end   int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initSegTree(l, r, node int) int {
	if l == r {
		segTree[node] = arr[l]
		return segTree[node]
	}
	mid := int((l + r) / 2)
	segTree[node] = min(initSegTree(l, mid, node*2), initSegTree(mid+1, r, node*2+1))
	return segTree[node]
}

func findMin(l, r, node, LB, RB int) int {
	if l > RB || r < LB {
		return 2000000000
	}
	if l >= LB && r <= RB {
		return segTree[node]
	}
	mid := int((l + r) / 2)
	return min(findMin(l, mid, node*2, LB, RB), findMin(mid+1, r, node*2+1, LB, RB))
}

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func Start() {
	sc := bufio.NewScanner(os.Stdin)
	r := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n := strToInt(sc.Text())
	sc.Scan()
	m := strToInt(sc.Text())
	arr = make([]int, n+1)
	queries := make([]query, m)
	segTree = make([]int, 4*n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(r, &arr[i])
	}
	for i := 0; i < m; i++ {
		q := query{}
		sc.Scan()
		a := strToInt(sc.Text())
		sc.Scan()
		b := strToInt(sc.Text())
		q.start = a
		q.end = b
		queries[i] = q
	}
	initSegTree(1, n, 1)
	for _, query := range queries {
		fmt.Fprintln(wr, findMin(1, n, 1, query.start, query.end))
	}
	wr.Flush()
	fmt.Println(segTree)
	fmt.Println(arr)
	fmt.Println(queries)
	fmt.Println(arr[2] + arr[1])
}
