package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"-"`
	// 포인터로 타입을 지정하기
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// 모의 DB
var courses []Course

// 미들웨어 헬퍼 -file
func (c *Course) IsEmpty() bool {
	// id와 강의명이 비어있는지 검증합니다.
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	// 순서와 두번째 인자는 포인터라는 것을 잘 외우도록 합니다.
	w.Write([]byte("<h1>API 생성</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("모든 강좌를 가져옵니다.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Course라고만 명명하는 것이 좋은 컨벤션입니다. 요청 유형을 굳이 작성하지 않습니다.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// 선형탐색으로 존재하는 강좌를 찾아냅니다.
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// 요청으로
	params := mux.Vars(r)

	// 응답을 강의로 돌려줍니다.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("강의가 없습니다.")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// 유효성 검증을 진행합니다. 비어있을 때를 대응합니다.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please and some data")
	}

	// {}이런 형식의 요청을 대응합니다.
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON에 데이터가 없습니다.")
		return
	}
	// uid와 문자열을 생성합니다. 데이터를 뒤에 붙입니다.

	// TODO: 중복 요청이 발생했는지 확인합니다.

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // 정수를 문자열로 변환합니다.
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			// id를 선택하고 덮어씁니다.
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	// TODO: 발견하지 못한 경우를 처리합니다.
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			break
		}
	}
}
