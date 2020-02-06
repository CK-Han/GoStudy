package main

import "fmt"

func main() { // 함수 및 조건문 등 모든 중괄호는 구문의 맨 뒤(다음 줄에 쓰면 컴파일 에러 ㄷㄷ)- 컴파일 시 줄의 마지막에 세미콜론을 자동으로 붙임
	// 들여쓰기는 탭
	i := 10

	if i >= 5 {
		fmt.Println("over 5") // 세미콜론은 자동으로 붙어 필요하지 않음, 여러줄을 쓰고 싶으면 사용은 가능
	}

	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}
