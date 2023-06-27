package main

import (
	"fmt"

	"time"
)

func main() {

	go start()

	fmt.Println("Started")

	time.Sleep(1 * time.Second)

	fmt.Println("Finished")

}

func start() {

	go start2()

	fmt.Println("In Goroutine")

}

func start2() {

	fmt.Println("In Goroutine2")

}

// در کد بالا داخل تابع main ما تابع start را با گوروتین اجرا کردیم و داخل تابع start تابع start2 را با گوروتین اجرا کردیم.
//این ۲ تابع start و start2 در کنار هم اجرا می شود و در نهایت کارشان اتمام می شود و هیچ کدام منتظر دیگری نخواهد بود.
