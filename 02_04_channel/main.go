package main

import (
	"fmt"

	"time"
)

func main() {

	ch := make(chan int)

	go send(ch)

	go receive(ch)

	time.Sleep(time.Second * 2)

}

func send(ch chan int) {

	time.Sleep(time.Second * 1)

	fmt.Println("Timeout finished")

	ch <- 1

}

func receive(ch chan int) {

	val := <-ch

	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)

}


// اتفاقی که در کد فوق رخداد زمانیکه شما عملیات دریافت را انجام می دهید تا زمانیکه مقداری از کانال
//دریافت نشود اون بخش از کد شما Lock می شود و پس از اینکه دریافت شد مقدار از کانال آن بخش Unlock خواهد شد.