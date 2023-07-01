package main

import (
	"context"

	"fmt"
)

func main() {

	ctx := context.WithValue(context.Background(), "language", "Go")

	fmt.Println(manager(ctx, "language"))

	fmt.Println(manager(ctx, "ahmad"))

}

func manager(ctx context.Context, key string) string {

	if v := ctx.Value(key); v != nil {

		return v.(string)

	}

	return "not found value"

}


// در کد فوق ما یک context ایجاد کردیم و داخلش با استفاده از WithValue مقدار key/value قرار دادیم و سپس این context را تابع manager
// پاس دادیم و داخل تابع manager ما با استفاده از متد Value که داخل اینترفیس ctx هست مقدار کلید language را گرفتیم.
