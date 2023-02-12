package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// 시드 생성
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("You can move 2")
	case 3:
		fmt.Println("You can move 3")
		fallthrough
	case 4:
		fmt.Println("You can move 4")
		fallthrough
	case 5:
		fmt.Println("You can move 5")
	case 6:
		fmt.Println("You can move 6")
	default:
		fmt.Println("What was that?")
	}
}
