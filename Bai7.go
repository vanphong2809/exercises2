package main

import (
	"context"
	"fmt"
	"time"
)

//Bai7 context
func Bai7(ctx context.Context, k string) {
	fmt.Println("Bai 7")
	before := ctx.Value(k).(int64)
	var now int64
	select {
	case <-time.After(3 * time.Second):
		now = time.Now().UnixNano()
	}
	fmt.Println("now-before", now-before)
}
