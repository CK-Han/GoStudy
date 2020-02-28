package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 고루틴 반환값은 쓸 수도, 쓰이지도 않음(무시됨)
func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

// Go 언어는 기본적으로 단일 코어를 사용
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "hello~"

	// 고루틴은 반복문이 모두 지나고 생성된다
	// 1 고루틴으로 i를 출력시 100이 된 후 고루틴이 시작
	for i := 0; i < 100; i++ {
		go func(n int) {
			//fmt.Println(s, n)
		}(i)

		go func() {
			fmt.Println(s, i)
		}()
	}

	fmt.Scanln()
}
