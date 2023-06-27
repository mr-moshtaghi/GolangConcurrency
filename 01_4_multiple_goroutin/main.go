package main

import (
	"fmt"

	"time"
)

func execute(id int) {

	fmt.Printf("id: %d\n", id)

}

func main() {

	fmt.Println("Started")

	for i := 0; i < 10; i++ {

		go execute(i)

	}

	time.Sleep(time.Second * 2)

	fmt.Println("Finished")

}

// در کد فوق ما یک حلقه قرار دادیم از i صفر تا ۱۰ که داخلش تابع execute را ۱۰ بار اجرا می کند.
//و هربار اجرا می شود خروجی های مختلفی دارد و ترتیبی درست نخواهید علت این اتفاق این است بطور همزمان اجرا می شوند در کنار
//هم و هرکدام از گوروتین ها زودتر کارش تمام شود خروجی را نمایش می دهد. به همین دلیل ترتیب درستی نخواهد داشت.
