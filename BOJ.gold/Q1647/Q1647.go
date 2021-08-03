package q1647

type node struct {
	now int
	to  int
	dis int
}

var N, M int
var parent = []int{}
var nodes = []node{}

func union(x, y int) {
	if parent[x] == parent[y] {
		return
	}
	if x > y {
		x, y = y, x
	}

	pX := findParent(x)
	pY := findParent(y)

	parent[pY] = parent[pX]

}

func findParent(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = findParent(parent[x])
	return parent[x]
}

func main() {
	// fmt.Scan(&N, &M)
	// parent = make([]int, N+1)
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)
	// for i := 0; i < M; i++ {
	// 	scanner.Scan()
	// 	now, _ := strconv.Atoi(scanner.Text())
	// 	scanner.Scan()
	// 	to, _ := strconv.Atoi(scanner.Text())
	// 	scanner.Scan()
	// 	dis, _ := strconv.Atoi(scanner.Text())
	// 	pq := make(PriorityQueue, 0)
	// }
	// for i := 1; i <= N; i++ {
	// 	parent[i] = i
	// }
}
