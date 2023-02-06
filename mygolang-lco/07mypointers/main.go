package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
	// 정수형을 할당할 포인터를 지정하는 방법입니다.
	// var ptr *int
	// fmt.Println("포인터의 값:", ptr)

	myNumber := 23
	// 포인터로 메모리상 동일한 23을 바라봅니다. &은 참조한다는 의미입니다.
	var ptr = &myNumber
	fmt.Println("포인터의 주소 값: ", ptr)
	fmt.Println("포인터의 실제 값: ", *ptr)
	*ptr = *ptr + 2

	fmt.Println("새로운 값:", myNumber)
}
