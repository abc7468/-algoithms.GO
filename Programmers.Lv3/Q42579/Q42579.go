package q42579

import (
	"container/heap"
)

type Song struct {
	num int
	cnt int
}

type songHeap []*Song

func (s songHeap) Len() int      { return len(s) }
func (s songHeap) Swap(a, b int) { s[a], s[b] = s[b], s[a] }
func (s *songHeap) Push(x interface{}) {
	*s = append(*s, x.(*Song))
}
func (s songHeap) Less(a, b int) bool {
	if s[a].cnt == s[b].cnt {
		return s[a].num < s[b].num
	}
	return s[a].cnt > s[b].cnt
}
func (s *songHeap) Pop() interface{} {
	older := *s
	n := len(older)
	val := older[n-1]
	*s = older[:n-1]
	older[n-1] = nil

	return val
}

type Genre struct {
	name   string
	amount int
}

type genreHeap []*Genre

func (g genreHeap) Len() int            { return len(g) }
func (g genreHeap) Less(a, b int) bool  { return g[a].amount > g[b].amount }
func (g genreHeap) Swap(a, b int)       { g[a], g[b] = g[b], g[a] }
func (g *genreHeap) Push(x interface{}) { *g = append(*g, x.(*Genre)) }
func (g *genreHeap) Pop() interface{} {
	older := *g
	n := len(older)
	val := older[n-1]
	*g = older[:n-1]
	older[n-1] = nil

	return val
}

func (g genreHeap) isEmpty() bool {
	if len(g) == 0 {
		return true
	}
	return false
}

func solution(genres []string, plays []int) []int {
	genrePlayCnt := make(map[string]*Genre)
	sMap := make(map[string]*songHeap)
	var g genreHeap

	for i, v := range genres {
		if _, ok := genrePlayCnt[v]; !ok {
			var s songHeap
			sMap[v] = &s
			heap.Push(sMap[v], &Song{i, plays[i]})
			genrePlayCnt[v] = &Genre{v, plays[i]}
		} else {
			genrePlayCnt[v].amount = genrePlayCnt[v].amount + plays[i]
			heap.Push(sMap[v], &Song{i, plays[i]})
		}
	}
	for _, v := range genrePlayCnt {
		g.Push(v)
	}
	size := len(g)
	sortedKeySet := make([]string, size)

	for i := 0; i < size; i++ {
		sortedKeySet[i] = heap.Pop(&g).(*Genre).name
	}
	answer := []int{}
	for _, v := range sortedKeySet {
		if sMap[v].Len() >= 2 {
			answer = append(answer, heap.Pop(sMap[v]).(*Song).num)
			answer = append(answer, heap.Pop(sMap[v]).(*Song).num)
		} else {
			answer = append(answer, heap.Pop(sMap[v]).(*Song).num)
		}

	}
	return answer
}

func Start() {
	genres := []string{"classic", "pop", "classic", "classic", "pop", "a", "hiphop"}
	plays := []int{500, 600, 150, 800, 2500, 800, 700}
	solution(genres, plays)
}
