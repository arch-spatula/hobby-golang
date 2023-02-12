package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")
	greeter()

	result := adder(3, 5)

	fmt.Println("간단한 결과", result)

	proResult, myMessage := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("간단한 결과", proResult, myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("안녕하세요")
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "adfsdsaf"
}
