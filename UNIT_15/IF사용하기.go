package main

import "fmt"

func main() {
	bValue := 1 // if의 조건값은 반드시 bool, 아니면 컴파일 오류
	_= bValue
	bV := false
	if bV {
		fmt.Println("true!")
	} else { // 이러한 else if 나 else 포맷
		fmt.Println("false!")
	}
}