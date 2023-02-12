package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	// 마지막 인자는 버퍼입니다.
	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}
	// golang의 화살표연산자는 할당합니다. 반대방향은 없습니다.
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// 받기 전용입니다.
		close(myCh) // 이렇게 되면 에러를 돌려줄 것입니다.
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// 보내기 전용입니다.
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
