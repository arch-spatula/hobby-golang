package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// log.Fatal로 에러처리
	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("mod 사용자")
}

// 백엔드의 흔한 문법입니다.
func serveHome(w http.ResponseWriter, r *http.Request) {
	// r은 요청이고 w는 응답입니다.
	w.Write([]byte("<h1>hello golang</h1>"))
}
