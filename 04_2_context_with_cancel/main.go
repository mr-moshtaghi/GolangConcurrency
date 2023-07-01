package main

import (
	"context"

	"fmt"

	"time"
)

func main() {

	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go task(cancelCtx)

	time.Sleep(time.Second * 3)

	cancelFunc()

	time.Sleep(time.Second * 1)

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

// در کد فوق ما یک context فرزند با استفاده از WithCancel ایجاد کردیم
//که به عنوان خروجی cancelCtx و cancelFunc را داد. سپس cancelCtx را به تابع task منتقل کردیم تا عملیاتی را انجام دهد.
//حال در ادامه کد تابع main ما یک Sleep در حد ۳ ثانیه گذاشتیم و گفتیم تابع cancelFunc اجرا شود. اگر دقت کنید
//پس ۳ ثانیه سیگنال لغو به تابع task ارسال شده و خطای Gracefully exit را چاپ کردیم و پس از آن خطای context چاپ کردیم.
