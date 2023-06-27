package main

import (
	"fmt"
)

func main() {

	go start()

	fmt.Println("Started")

	fmt.Println("Finished")

}

func start() {

	fmt.Println("In Goroutine")

}

// در خروجی بالا هرگز پیغام داخل تابع ()start چاپ نمی شود.
// علت اصلی این اتفاق این است که تابع main خودش داخل یک گوروتین اجرا می شود و زمانیکه شما یک تابع دیگری را داخل
// گوروتین قرار می دهید تا لحظه ای که تابع برای اجرا برنامه ریزی شود برنامه اتمام می شود.
