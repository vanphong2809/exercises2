package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//bai 1
	//Cach 1
	// Bai1()
	Bai1C2()

	//Bai 2
	Bai2()

	//Bai4
	Bai4(1592190294764144364)

	//Bai5
	Bai5()

	//Bai7
	k := "key"
	ctx := context.WithValue(context.Background(), k, time.Now().UnixNano())
	fmt.Println("ctx", ctx.Value(k))
	Bai7(ctx, k)

	//Bai8
	Bai8()
	//Bai 9
	Bai9()
}
