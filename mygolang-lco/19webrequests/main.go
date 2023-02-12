package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web에 요청합니다.")
	// get 요청을 날립니다.
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response 타입: %T\n", response)

	// 요청이 종료되면 닫도록 합니다.
	defer response.Body.Close()

	dateBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dateBytes)
	fmt.Println(content)
}
