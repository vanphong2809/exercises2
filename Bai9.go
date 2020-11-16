package main

import (
	"fmt"
	"time"
)

//Bai9 Tạo 1 đoạn code sử dụng `time.AfterFunc()`, sau 100ms thì in ra dòng chữ `i'm study`
func Bai9() {
	fmt.Println("Bai 9: ")
	durationOfTime := time.Duration(100) * time.Millisecond
	f := func() {
		fmt.Println("i am study")
	}
	Timer := time.AfterFunc(durationOfTime, f)
	defer Timer.Stop()
	time.Sleep(2 * time.Second)
}
