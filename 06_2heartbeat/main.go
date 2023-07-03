package main

import (
	"fmt"

	"time"
)

func heartbeat(interval time.Duration, c chan<- struct{}) {

	ticker := time.NewTicker(interval)

	for {

		select {

		case <-ticker.C:

			c <- struct{}{}

		}

	}

}

func task() {

	// Do some work here

	fmt.Println("Task running...")

}

func main() {

	c := make(chan struct{})

	go heartbeat(1*time.Second, c)

	for {

		select {

		case <-c:

			task()

		}

	}

}

// در کد فوق ما یک تابع به نام heartbeat ایجاد کردیم که طی مدت زمانی یک سیگنال می فرستد تا تابع task اجرا شود که وضعیت فرآیند یا سرویس
//را گزارش دهد. ما مدت زمان را ۱ ثانیه گذاشتیم و یک کانال ساختار ایجاد کردیم و به تابع heartbeat که داخل یک گوروتین هست
//پاس دادیم سپس هر ۱ ثانیه از طریق کانال c ما سیگنال اجرای task برای بررسی وضعیت سرویس یا فرآیند را دریافت میکنیم.
