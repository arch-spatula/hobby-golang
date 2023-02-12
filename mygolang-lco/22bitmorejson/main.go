package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json: "coursename"`
	Price    int
	Platform string   `json: "website"`
	password string   `json: "-"` // 제거
	Tags     []string `json: "tags, omitempty"`
}

func main() {
	fmt.Println("JSON입니다.")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS BootCamp", 299, "LearnCodeOnline.in", "asdf", []string{"web-dev", "js"}},
		{"MERN BootCamp", 199, "LearnCodeOnline.in", "qwer", []string{"full-stack", "js"}},
		{"Vue BootCamp", 299, "LearnCodeOnline.in", "zcvx", nil},
	}

	// json데이터로 패키징합니다.
	// 첫번째 인자는 인터페이스입니다. struct의 파생된 단어입니다.
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "ReactJS BootCamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Tags": [
						"web-dev",
						"js"
		]
}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON은 검증되었습니다.")
		// 두번째 인자에 참조할 주소를 전달합니다.
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON은 유효하지 않습니다.")
	}

	// 키와 값같은 해쉬테이블 자료형 응용
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
