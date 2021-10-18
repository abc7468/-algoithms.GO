package test2

import "fmt"

var holidayMap map[int]interface{}

func checkDay(day string) {
	switch day {
	case "MON":
		addHoliday(6, 0)
	case "TUE":
		addHoliday(5, 6)

	case "WED":
		addHoliday(4, 5)

	case "THU":
		addHoliday(3, 4)

	case "FRI":
		addHoliday(2, 3)

	case "SAT":
		addHoliday(1, 2)

	case "SUN":
		addHoliday(0, 1)
	}
}

func addHoliday(a, b int) {
	for i := 0; i <= 30; i++ {
		if i%7 == a || i%7 == b {
			holidayMap[i] = nil
		}
	}
}

func solution(leave int, day string, holidays []int) int {
	holidayMap = make(map[int]interface{})
	checkDay(day)
	for _, v := range holidays {
		holidayMap[v] = nil
	}
	start := 1
	cnt := leave
	max := 0
	for end := 1; end <= 30; end++ {
		if _, ok := holidayMap[end]; !ok {
			cnt--
		}
		if cnt < 0 {
			if _, ok := holidayMap[start]; !ok {
				cnt++
			}
			start++

		}
		fmt.Println(start, end, cnt)

		if max < end-start+1 {
			max = end - start + 1
		}
	}
	fmt.Println(max)
	return max
}
func Start() {
	leave := 4
	day := "FRI"
	holidays := []int{6, 21, 23, 27, 28}
	solution(leave, day, holidays)
}
