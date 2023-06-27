//package main
//
//func main() {
//
//	ch := make(chan int)
//
//	close(ch)
//
//	ch <- 2
//
//}

package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 1)

	ch <- 2

	val, ok := <-ch

	fmt.Printf("Val: %d OK: %t\n", val, ok)

	close(ch)

	val, ok = <-ch

	fmt.Printf("Val: %d OK: %t\n", val, ok)

}
