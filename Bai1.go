package main

import (
	"fmt"
	"time"
)

//Bai1  Tạo 1 chương trình cứ 3s in ra dòng chữ: `time now: {milliseconds}`:
//Sau khi in được 3 lần thì in ra dòng chữ `kết thúc`
func Bai1() {
	fmt.Println()
	fmt.Println("Bai 1 cach 1")
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("time now: ", time.Now().UnixNano()/int64(time.Millisecond))
	}
	fmt.Println("kết thúc")
}

//Bai1C2 Tạo 1 chương trình cứ 3s in ra dòng chữ: `time now: {milliseconds}`:
//Sau khi in được 3 lần thì in ra dòng chữ `kết thúc`
func Bai1C2() {
	fmt.Println()
	fmt.Println("Bai 1 cach 2")
	ticker := time.NewTicker(3 * time.Second)
	myChannel := make(chan bool)
	go func() {
		for {
			select {
			case <-myChannel:
				return
			case t := <-ticker.C:
				{
					fmt.Println("time now: ", t)
				}
			}
		}
	}()
	time.Sleep(9100 * time.Millisecond)
	ticker.Stop()
	fmt.Println("test")
	myChannel <- true
	fmt.Println("Kết thúc")
}
