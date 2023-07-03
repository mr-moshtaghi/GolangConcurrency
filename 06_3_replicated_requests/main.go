package main

import (
	"fmt"

	"net/http"
)

func makeRequest(url string, c chan<- *http.Response) {

	resp, err := http.Get(url)

	if err != nil {

		c <- nil

	} else {

		c <- resp

	}

}

func main() {

	urls := []string{"http://example.com", "http://example.org", "http://example.net"}

	c := make(chan *http.Response)

	defer close(c)

	for _, url := range urls {

		go makeRequest(url, c)

	}

	for i := 0; i < len(urls); i++ {

		resp := <-c

		if resp == nil {

			fmt.Println("Error making request")

		} else {

			fmt.Println(resp.Status)

		}

	}

}

// در کد فوق ما یک تابع makeRequest داریم که ۲ تا پارامتر ورودی دارد اولین پارامتر
//url میگیرد و دومین پارامتر یک کانال فقط ارسال از نوع http.Response* میگیرد.
//سپس یک ریکوئست با متد GET ایجاد میکند و خروجی را داخل کانال میفرستد. در تابع main ما یک لیست url داریم که قرار است
//بصورت موازی به این آدرس ها درخواست بفرستیم و خروجی را دریافت کنیم در اینجا یک کانال از نوع http.Response* ایجاد
//کردیم و سپس یک حلقه for-range قرار دادیم به ازای هر یک از url ها تابع makeRequest را فراخوانی کردیم و داخل گوروتین
//قرار دادیم. در نهایت یک حلقه for-i داریم که به تعداد url ها شمارش میکند و از طریق کانال رسپانس را دریافت می کند.
