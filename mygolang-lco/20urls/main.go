package main

import (
	"fmt"
	"net/url"
)

// 가짜 url을 만들어봅니다.

const MY_URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid:afsdqwerzxv3645"

func main() {
	fmt.Println("URL 다루기에 환영합니다.")
	fmt.Println(MY_URL)

	// 파싱 처리
	result, _ := url.Parse(MY_URL)
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid:afsdqwerzxv3645
	fmt.Println(result.Port())   // 3000

	qparams := result.Query()
	fmt.Printf("Query Param: %T\n", qparams) // Query Param: url.Values

	fmt.Println(qparams["coursename"]) // [reactjs]

	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	// 포인터를 전달합니다.
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jake",
	}

	// 문자열로 변환하고 할당합니다.
	auotherURL := partsOfUrl.String()
	fmt.Println(auotherURL)
}
