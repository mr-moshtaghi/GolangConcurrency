package main

import (
	"fmt"

	"time"
)

func main() {

	ch := make(chan int)

	go sum(ch, 3)

	ch <- 2

	ch <- 2

	ch <- 2

	close(ch)

	time.Sleep(time.Second * 1)

}

func sum(ch chan int, len int) {

	sum := 0

	for i := 0; i < len; i++ {

		sum += <-ch

	}

	fmt.Printf("Sum: %d\n", sum)

}
