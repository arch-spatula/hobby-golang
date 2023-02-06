package main

import "fmt"

func main() {
	fmt.Println("Structs에 온것을 환영합니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30}
	fmt.Println(jake)
	fmt.Printf("Jake details are: %+v\n", jake)
	fmt.Printf("Name is %v and email is %v.", jake.Name, jake.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
