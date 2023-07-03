package main

import (
	"fmt"

	"time"
)

func task1(done chan struct{}) {

	// Do some work here

	fmt.Println("doing task 1")

	time.Sleep(2 * time.Second)

	fmt.Println("task 1 has been completed")

	done <- struct{}{}

}

func task2(done <-chan struct{}) {

	select {

	case <-done:

		// Do some work here

		fmt.Println("doing task 2")

		time.Sleep(2 * time.Second)

		fmt.Println("task 2 has been completed")

	}

}

func main() {

	done := make(chan struct{})

	go task1(done)

	go task2(done)

	time.Sleep(5 * time.Second)

	fmt.Println("all tasks has been completed")

}

// در کد فوق ما دو تابع به نام task1 و task2 داریم که داخل گوروتین هستند و به عنوان پارامتر ورودی
//یک کانال ساختار خالی میگیرد حال کانال ساختار را به هر دو تابع پاس دادیم در ابتدا task1 فرآیند خود را انجام می دهد
//و سپس سیگنال انجام شدن را می فرستد و در ادامه تابع task2 سیگنال را دریافت می کند و فرآیندهای خود را انجام می دهد.
