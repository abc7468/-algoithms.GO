package test3

import "fmt"

func solution(v [][]int) []int {
    answer := []int{}
    xCnt map[int]int := make(map[int]int)
    yCnt map[int]int := make(map[int]int)
    for _, val := range v{
        if _, ok := xCnt[val[0]]; !ok{
            xCnt[val[0]] = 1
        }else{
            xCnt[val[0]] = xCnt[val[0]]+1
        }
        if _, ok := xCnt[val[1]]; !ok{
            xCnt[val[1]] = 1
        }else{
            xCnt[val[1]] = xCnt[val[1]]+1
        }
    }
    for key, val := range xCnt{
        if val == 1{
            answer = append(answer, key)
        }
    }
    for key, val := range yCnt{
        if val == 1{
            answer = append(answer, key)
        }
    }
    return answer
}