package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	fmt.Printf("Capacity: %d\n", cap(ch))

}

//     توجه کنید ظرفیت کانال بافر نشده همیشه صفر است.
