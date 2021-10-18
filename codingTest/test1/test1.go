package test1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var sList []string
var nList []int

func findNumber(array []int, start int) int {
	index := 0
	for index = 0; index < len(array); index++ {
		if array[index] == start {
			break
		}
	}
	for i := index + 1; i < len(array); i++ {
		start++
		if array[i] == start {
			continue
		}
		return start
	}
	return start + 1
}
func findNewId(registered_list []string, new_id string) bool {
	for _, str := range registered_list {
		if str == new_id {
			return true
		}
	}
	return false
}

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func solution(registered_list []string, new_id string) string {
	if !findNewId(registered_list, new_id) {
		return new_id
	}

	var index int
	for i := 0; i < len(new_id); i++ {
		if new_id[i] >= '0' && new_id[i] <= '9' {
			index = i
			break
		}
		index = len(new_id)
	}

	newIdS := new_id[:index]
	newIdN := strToInt(new_id[index:])
	r, _ := regexp.Compile("^[a-z]{3,6}")
	for _, str := range registered_list {
		endIdx := r.FindStringIndex(str)[1]
		if index == endIdx {
			if newIdS == str[:endIdx] {
				nList = append(nList, strToInt(str[endIdx:]))
			}
		}
	}
	sort.Ints(nList)
	numVal := findNumber(nList, newIdN)
	return newIdS + fmt.Sprint(numVal)
}
func Start() {
	registry_list := []string{"card", "card1", "card2", "card3", "card5"}

	new_id := "card5"

	fmt.Println(solution(registry_list, new_id))
}
