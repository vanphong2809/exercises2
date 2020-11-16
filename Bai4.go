package main

import "fmt"

//Bai4 Tính ra số phút(number_of_minutes) của mốc thời gian sau `1592190294764144364`
func Bai4(nanosecond int64) {
	fmt.Println()
	fmt.Println("Bai4")
	numberofminutes := (nanosecond / (60 * 1000000000))
	fmt.Println("So phut sau khi quy doi: ",numberofminutes)
}
