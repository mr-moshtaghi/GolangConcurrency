package main

import (
	"context"

	"fmt"

	"time"
)

func unhealthyGoroutine(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("Goroutine is unhealthy, exiting")

			return

		default:

			// Do some work here

			fmt.Println("Goroutine running...")

			time.Sleep(500 * time.Millisecond)

		}

	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	for {

		go unhealthyGoroutine(ctx)

		<-time.After(4 * time.Second)

	}

}

// در کد فوق ما یک تابع به نام unhealthyGoroutine داریم که بصورت جداگانه در گوروتین های مختلف اجرا می شود و کاری را انجام می دهد.
//داخل تابع ما یک select داریم که در یکی از case هایش context.Done را بررسی میکنیم آیا فرآیند لغو شده است یا خیر.
//داخل تابع main ما یک context از نوع Timeout با مدت زمان ۳ ثانیه ای ایجاد کردیم و در ادامه داخل یک حلقه بینهایت
//تابع unhealthyGoroutine داخل گوروتین قرار دادیم و هر ۴ ثانیه یک نمونه از این تابع داخل گوروتین های مختلف اجرا می شود.

// در اینجا کارهای داخل تابع unhealthyGoroutine انجام شود پس از ۳ ثانیه بواسطه context فرآیندها لغو می شود و از گوروتین
//خارج می شود. حال ما داخل تابع main اجازه دادیم یک گوروتین جدید و سالم را اجرا کند و جایگزین گوروتین ناسالم شود.
