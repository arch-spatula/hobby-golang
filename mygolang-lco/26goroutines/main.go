package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// 모두 각각 포인터로 제어해야 합니다.
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		// 비동기 작업을 시작합니다.
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// 처리가 완료된 것을 마지막에 처리합니다.
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println("%d 200 status code for %s ", res.StatusCode, endpoint)
	}
}
