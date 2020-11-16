package main

import (
	"context"
	"fmt"
	"time"
)

//Bai3 Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec
//Nhưng sau 3s thì kết thúc hàm đấy
//Sử dụng và tìm hiểu context. Nêu được tác dụng của context trong chương trình.
func Bai3() {
	fmt.Println("Bai3")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(3 * time.Second)
			fmt.Println("Lan ", i)
		}
	}
}
