package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok 문법입니다. err ok 문법이라고 합니다.
	// golang은 try catch가 없습니다. 에러를 부울처럼 취급하기 바라면서 설계했습니다.

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this is %T", input)
}
