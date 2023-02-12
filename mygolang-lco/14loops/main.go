package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, value)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		// 종료
		if rougeValue == 5 {
			break
		}

		// 하나 건너 뛰기
		if rougeValue == 3 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto loc
		}

		fmt.Println("값: ", rougeValue)
		rougeValue++
	}

loc:
	fmt.Println("go to 기능을 구현합니다.")
}
