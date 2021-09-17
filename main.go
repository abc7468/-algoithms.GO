package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go receiver(ch)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	fmt.Println("Finish!")
}

func receiver(ch <-chan int) {
	for i := range ch {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
}
