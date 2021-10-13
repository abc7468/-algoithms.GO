package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var str1, str2 string
	fmt.Fscan(r, &str1, &str2)
	p := 1
	str1Size := len(str1)
	str2Size := len(str2)
	comp := 0

	for i := 0; i < str2Size; i++ {
		comp = comp<<1 + int(str2[i])
	}
	p = 1
	val := 0
	for i := 0; i < str2Size; i++ {
		val = val<<1 + int(str1[i])
		p = p << 1
	}

	p = p >> 1
	tmp := 0
	if val == comp {
		for i := 0; i < str2Size; i++ {
			if str1[i] != str2[i] {
				break
			}
			tmp++
		}
		if tmp == str2Size {
			fmt.Println(1)
			return
		}
	}
	for i := 1; i < str1Size-str2Size+1; i++ {
		val -= int(str1[i-1]) * p
		val = val << 1

		val += int(str1[i+str2Size-1])
		tmp := 0
		if val == comp {
			for j, v := range str2 {
				if v != rune(str1[i+j]) {
					break
				}
				tmp++
			}
			if tmp == str2Size {
				fmt.Println(1)
				return
			}
		}
	}

	fmt.Println(0)
}
