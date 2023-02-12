package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("요청처리 1장입니다.")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const MY_URL = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "jake")
	data.Add("lastName", "the dog")
	data.Add("email", "jakethegod@go.dev")

	response, err := http.PostForm(MY_URL, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

// func PerformPostJsonRequest() {
// 	const MY_URL = "http://localhost:8000/post"

// 	requestBody := strings.NewReader(`
// 		{
// 			"courseName": "Let's go with golang",
// 			"price": "0",
// 			"platform": "learnCodeOnline.in"
// 		}
// 	`)

// 	// 두번째 인자 header의 content-type을 주의하도록 합니다.
// 	response, err := http.Post(MY_URL, "application/json", requestBody)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()

// 	content, _ := ioutil.ReadAll(response.Body)

// 	fmt.Println(string(content))
// }

// public 함수로 만듭니다.
// func PerformGetRequest() {
// 	const MY_URL = "http://localhost:8000/get"

// 	response, err := http.Get(MY_URL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()
// 	fmt.Println("상태 코드: ", response.StatusCode)
// 	fmt.Println("콘텐츠 길이: ", response.ContentLength)

// 	// fmt.Println(string(content))

// 	var responseString strings.Builder
// 	content, _ := ioutil.ReadAll(response.Body)
// 	byteCount, _ := responseString.Write(content)
// 	fmt.Println("byteCount: ", byteCount)
// 	fmt.Println("byteCount: ", responseString.String()) // 장점은 원본 데이터를 달고 있습니다.

// }
