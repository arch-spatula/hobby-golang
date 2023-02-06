package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang if else")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "일반 유저"
	} else if loginCount > 10 {
		result = "해비유저"
	} else {
		result = "이벤트 달성"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("짝수")
	} else {
		fmt.Println("홀수")
	}

	// 요청을 받으면서 할당하게 만드는 방법입니다.
	if num := 3; num < 10 {
		fmt.Println("num은 10보다 큽니다.")
	} else {
		fmt.Println("num은 10보다 작습니다.")
	}

}
