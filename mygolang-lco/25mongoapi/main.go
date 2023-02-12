package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arch-spatula/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("서버 시작")

	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("3000에 듣는 중")

}
