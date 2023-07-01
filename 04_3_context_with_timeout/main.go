package main

import (
	"context"

	"fmt"

	"time"
)

func main() {

	ctx := context.Background()

	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()

	go task1(cancelCtx)

	time.Sleep(time.Second * 4)

}

func task1(ctx context.Context) {

	i := 1

	for {

		select {

		case <-ctx.Done():

			fmt.Println("Gracefully exit")

			fmt.Println(ctx.Err())

			return

		default:

			fmt.Println(i)

			time.Sleep(time.Second * 1)

			i++

		}

	}

}

// در کد فوق ما یک context فرزند با استفاده از تابع WithTimeout ایجاد کردیم و مدت زمان ۳ ثانیه
//به این تابع پاس دادیم و پس از آن context فرزند به همراه تابع cancelFunc دریافت کردیم. حالا تابع cancel را
//داخل defer قرار دادیم و cancelCtx را به تابع task1 که داخل گوروتین است پاس دادیم. و یک Sleep به مدت ۴ ثانیه گذاشتیم
//تابع main کارش اتمام نشود. حال پس از اینکه ۳ ثانیه گذشت داخل select سیگنال cancel را دریافت کردیم و خطای
//context deadline exceeded که نشان دهنده اتمام شدن مدت زمان هست را چاپ کردیم. و همانطور که متوجه شدید درخواست کلی ما لغو شدش.
