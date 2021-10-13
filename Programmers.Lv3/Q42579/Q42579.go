package q42579

import (
	"container/heap"
	"fmt"
)

type Genre struct {
	name   string
	amount int
}

type Song struct {
	num int
	cnt int
}

type sPQ []*Song

func (s sPQ) Len() int {
	return len(s)
}

func (s sPQ) Less(a, b int) bool {
	if s[a].cnt == s[b].cnt {
		return s[a].num < s[b].num
	}
	return s[a].cnt > s[b].cnt
}

func (s sPQ) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s *sPQ) Push(x interface{}) {
	item := x.(*Song)
	*s = append(*s, item)
}

func (s *sPQ) Pop() interface{} {
	older := *s
	n := len(older)
	val := older[n-1]
	*s = older[:n-1]
	older[n-1] = nil

	return val
}

type gPQ []*Genre

func (g gPQ) Len() int {
	return len(g)
}

func (g gPQ) Less(a, b int) bool {
	return g[a].amount > g[b].amount
}

func (g gPQ) Swap(a, b int) {
	g[a], g[b] = g[b], g[a]
}

func (g *gPQ) Push(x interface{}) {
	item := x.(*Genre)
	*g = append(*g, item)
}

func (g *gPQ) Pop() interface{} {
	older := *g
	n := len(older)
	val := older[n-1]
	*g = older[:n-1]
	older[n-1] = nil

	return val
}

func (g gPQ) isEmpty() bool {
	if len(g) == 0 {
		return true
	}
	return false
}

func solution(genres []string, plays []int) []int {
	genrePlayCnt := make(map[string]*Genre)
	sMap := make(map[string]*sPQ)
	var g gPQ

	for i, v := range genres {
		if _, ok := genrePlayCnt[v]; !ok {
			var s sPQ
			heap.Init(&s)
			sMap[v] = &s
			heap.Push(sMap[v], &Song{
				i,
				plays[i],
			})
			genrePlayCnt[v] = &Genre{
				name:   v,
				amount: plays[i],
			}
		} else {
			genrePlayCnt[v].amount = genrePlayCnt[v].amount + plays[i]
			heap.Push(sMap[v], &Song{
				i,
				plays[i],
			})
		}
	}
	for _, v := range genrePlayCnt {
		g.Push(v)
	}
	heap.Init(&g)
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
	fmt.Println(answer)

	return []int{}
}

func Start() {
	genres := []string{"classic", "pop", "classic", "classic", "pop", "a", "hiphop"}
	plays := []int{500, 600, 150, 800, 2500, 800, 700}
	solution(genres, plays)
}
