package q86051

func solution(numbers []int) int {
	answer := 0
	for i := 1; i < 10; i++ {
		cnt := 0
		for j := 0; j < len(numbers); j++ {
			if numbers[j] == i {
				break
			} else {
				cnt++
			}
		}
		if cnt == len(numbers) {
			answer += i
		}

	}
	return answer
}

func Start() {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 0}
	solution(numbers)
}
