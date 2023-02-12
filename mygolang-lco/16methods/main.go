package main

import "fmt"

func main() {
	fmt.Println("메서드를 다룹니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30, 23}
	// fmt.Println(jake)
	// fmt.Printf("Jake details are: %+v\n", jake)
	// fmt.Printf("Name is %v and email is %v.", jake.Name, jake.Email)
	jake.GetStatus()
	jake.NewMail()
	fmt.Println(jake)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int // 비공개 속성값은 케멀케이스로 정의합니다.
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "jakethedog@go.dev"
	fmt.Println("Email: ", u.Email)
}
