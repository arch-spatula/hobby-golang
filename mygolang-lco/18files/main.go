package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("golang으로 파일처리")
	content := "파일을 읽어볼 것입니다."

	// 파일 생성
	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	// 생성한 파일에 글 작성하기
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	// 생성한 파일을 닫기합니다.
	defer file.Close()
	readFile("./mylcogofile.txt")
}

// 파일 읽기는 생성과 유사합니다.

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("dataByte의 생김새는...", dataByte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
