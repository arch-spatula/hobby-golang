package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math")
	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("합: ", myNumberOne+int(myNumberTwo)) // 자동완성으로 형변환이 발생합니다.

	// 무작위 숫자를 만드는 법입니다.
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//  crypto를 통한 난수를 생성합니다.
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
