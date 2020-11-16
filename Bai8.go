package main

import (
	"fmt"
	"time"
)

//Bai8 Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ `${time.Now().Unix()} done`
func Bai8() {
	fmt.Println("Bai 8: ")
	for i := 0; i < 100; i++ {
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Println(time.Now().Unix(), " done")
		}
	}
}
