package q6459

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() int {
	if s.length() == 0 {
		return -1
	}
	older := *s
	val := older[len(older)-1]
	*s = older[:len(older)-1]
	return val
}
func (s stack) top() int {
	return s[len(s)-1]
}
func (s stack) length() int {
	return len(s)
}
func (s stack) isEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s *stack) clear() {
	for !s.isEmpty() {
		s.pop()
	}
}
func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}
func Start() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for {
		max := 0
		sc.Scan()
		input := strToInt(sc.Text())
		if input == 0 {
			return
		}
		arr := make([]int, input)
		for i := 0; i < input; i++ {
			sc.Scan()
			arr[i] = strToInt(sc.Text())
		}
		var s stack
		size := input
		left := make([]int, size)
		right := make([]int, size)
		s.push(0)
		left[0] = -1
		for i := 1; i < size; i++ {
			for arr[s.top()] >= arr[i] {
				s.pop()
				if s.isEmpty() {
					left[i] = -1
					s.push(i)
					break
				}
			}
			if s.top() == i {
				continue
			}
			left[i] = s.top()
			s.push(i)
		}
		s.clear()
		s.push(size - 1)
		right[size-1] = size
		for i := size - 2; i >= 0; i-- {
			for arr[s.top()] >= arr[i] {
				s.pop()
				if s.isEmpty() {
					right[i] = size
					s.push(i)
					break
				}
			}
			if s.top() == i {
				continue
			}
			right[i] = s.top()
			s.push(i)
		}
		for i := 0; i < size; i++ {
			if (right[i]-left[i]-1)*arr[i] > max {
				max = (right[i] - left[i] - 1) * arr[i]
			}
		}
		fmt.Println(max)
	}
}
