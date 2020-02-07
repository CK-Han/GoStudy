package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Int() % 10

	// GO 는 break 를 따로 표기 안해도 된다
	// 대신 강제적으로 break 표기하여 해당 case 스코프를 빠져나오기 가능
	// 기존 switch 문 처럼 쭉 내려가고싶으면 fallthrough
	// 혹은 case 한 줄에 같이쓰거나
	switch i {
	case 1:
		fmt.Println("hi 1")
		fallthrough
	case 2:
		fmt.Println("hi 2")
	case 3,4,5:
		fmt.Println("hi 3 4 5 ")
	default:
		if j := rand.Int() % 3; j > 1 {
			fmt.Println("bigger than 1")
			break
		}
		fmt.Println("hi")
		//fallthrough 맨 마지막은 불가
	}

	// GO switch 는 조건식도 가능하다 ㄷㄷ; 변수가 생략되는 형태
	j := 32
	switch {
	case j <= 16:
		fmt.Println("under 16")
	case j > 16:
		fmt.Println("over 16")
	}

	// GO switch 는 함수의 결과값으로도 가능
	switch n := rand.Uint64(); {
	case math.MaxInt64 < n:
		fmt.Println("bigger than int64")
	default:
		fmt.Println("between 0 and max(int64)")
	}
}
