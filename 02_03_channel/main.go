package main

import (
	"fmt"

	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel")

	go send(ch)

	fmt.Println("Receiving from channel")

	go receive(ch)

	time.Sleep(time.Second * 1)

}

func send(ch chan int) {

	ch <- 1

}

func receive(ch chan int) {

	val := <-ch

	fmt.Printf("Value Received=%d in receive function\n", val)

}

// در انتهای تابع main ما یک sleep به مدت ۱ ثانیه قرار دادیم که بتوانید
//خروجی برنامه را ببینیم و اگر اینکار را نکنیم برنامه متوقف می شود و ممکن دو گوروتین برای اجرا برنامه ریزی نشوند.
