package main

import (
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
func tinhsonamnhuan() int {
	count := 0
	for i := 1970; i < time.Now().Year(); i++ {
		if isLeapYear(i) {
			count++
		}
	}
	return count
}

//Bai2 Việt 1 đoạn chương trình tính ra ngày hiện tại theo timestamp.
//Điểm mốc từ mức 0  tại 1/1/1970
func Bai2() {
	fmt.Println()
	fmt.Println("Bai2")
	timestamp := ((time.Now().Year()-1970)*365+tinhsonamnhuan()+time.Now().YearDay()-1)*60*60*24*1000 + (time.Now().Hour()-7)*60*60*1000 + time.Now().Minute()*60*1000 + time.Now().Second()*1000
	fmt.Println(timestamp)
}
