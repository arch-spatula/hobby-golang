package main

import "fmt"

func main() {
	fmt.Println("Go의 배열에 온 것을 환영합니다.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("과일 목록입니다: ", fruitList)
	fmt.Println("과일 목록입니다: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("야채 목록:", len(vegList))
}
