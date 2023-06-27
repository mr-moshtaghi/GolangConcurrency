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

	fmt.Println("In Goroutine")

}

// در کد فوق ما تابع ()start را بواسطه کلمه کلیدی go داخل گوروتین قرار دادیم
//و این تابع بصورت مستقلانه از تابع main اجرا شد. اما این وسط یک نکته ای وجود دارد. همانطور که گفتیم تابع اصلی جهت
//اجرا برنامه های زبان گو تابع main می باشد و اگر شما تابعی را بواسطه گوروتین از main جدا کنید ممکن است فرآیندهای داخل
//تابع main زود اتمام شود و شما خروجی تابعی که داخل گوروتین گذاشتید را نبینید.
