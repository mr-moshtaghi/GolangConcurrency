package main


import "fmt"


func main() {

	ch1 := make(chan string)

	select {

	case msg := <-ch1:

		fmt.Println(msg)

	}

}

// اتفاقی که در کد فوق رخ داد ما یک کانال ایجاد کردیم و سپس داخل select یک case قرار دادیم
//که منتظر دریافت داده از کانال می باشد. اما چون هیچ داده به کانال
//ارسال نمی شود برنامه بطور کلی در همان تیکه از کد بلاک می شود و در نهایت شما با خطای داخل خروجی مواجه خواهید شد.