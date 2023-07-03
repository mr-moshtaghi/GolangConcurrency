package main

import (
	"fmt"
)

func fibo(n int) <-chan int {

	result := make(chan int)

	go func() {

		defer close(result)

		if n <= 2 {

			result <- 1

			return

		}

		result <- <-fibo(n-1) + <-fibo(n-2)

	}()

	return result

}

func main() {

	fmt.Println(<-fibo(25))

}
