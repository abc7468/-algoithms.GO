package q20291

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Start() {
	r := bufio.NewReader(os.Stdin)
	dic := make(map[string]int)
	var n int
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		tmp := ""
		fmt.Fscan(r, &tmp)

		ex := strings.Split(tmp, ".")[1]
		if _, ok := dic[ex]; !ok {
			dic[ex] = 1
		} else {
			dic[ex] = dic[ex] + 1
		}
	}
	var keys []string
	for k, _ := range dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s %d\n", key, dic[key])
	}
}
