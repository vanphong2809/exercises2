package main

import (
	"fmt"
	"time"
)

//Bai5 Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385`
func Bai5() {

	fmt.Println()
	fmt.Println("Bai5")
	day := time.Time(time.Unix(0, 1592190385000*int64(time.Millisecond)))
	fmt.Println("Ngay trong tuan o moc thoi gian 1592190385: ", day.Weekday())
	fmt.Println("Ngay trong tuan o moc thoi gian 1592190385: ", int(day.Weekday()))
}
