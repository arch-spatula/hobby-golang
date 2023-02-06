package main

import (
	"fmt"
)

func main() {
	fmt.Println("슬라이스에 왔습니다.")
	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 456
	// highScores[3] = 867
	// // highScores[4] = 777 // 오버플로우 오류가 발생합니다.

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// 인덱스를 기준으로 slice를 제어하는 법입니다.

	var courses = []string{"react.js", "javascript", "swift", "python", "php"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
