package main


import (

	"fmt"

)


func main() {

	ch := make(chan int, 1)

	ch <- 1

	ch <- 2

	fmt.Println("Sending value to channnel complete")

	val := <-ch

	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)

}

// کانال ما به خاطر پر شدن ظرفیتش بلاک شده بود و داده دیگه ای را نمی توانست
//نگه داری کند. در نتیجه با خطای deadlock مواجه شد و برنامه کاملا متوقف شد