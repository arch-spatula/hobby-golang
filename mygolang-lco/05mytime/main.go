package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 02-05-2023 03:02:21 Sunday

	// 날짜 생성
	createdDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
