package main

import (
	"context"

	"fmt"

	"time"
)

func main() {

	ctx := context.Background()

	cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))

	defer cancel()

	go task(cancelCtx)

	time.Sleep(time.Second * 6)

}

func task(ctx context.Context) {

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


// در کد فوق یک context فرزند با استفاده از تابع WithDeadline ایجاد کردیم و سپس با توجه
//به زمان فعلی مدت زمان ۵ ثانیه بعد را درنظر گرفتیم که مثلا اگر الان ساعت است 10:45:30 درخواست را در 10:45:35 لغو کند.
