package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("모든 언어 목록:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("모든 언어 목록:", languages)

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)

	}

}
