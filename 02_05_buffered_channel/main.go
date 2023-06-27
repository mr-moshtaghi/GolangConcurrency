package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 1)

	ch <- 1

	fmt.Println("Sending value to channnel complete")

	val := <-ch

	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)

}
