package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := 3, 2
	fmt.Printf("%016b\n", a & b)

	c, d := 255, 68 // 11111111, 01000100
	fmt.Printf("%08b\n", c &^ d) // AND NOT

	e := +5
	f := +e
	fmt.Println(f)

	// 비교 연산자의 동작과 주의점
	// 실수는 epsilon 주의
	// 문자열은 내용 같은지 확인
	// 배열은 요소 같은지 확인
	// 슬라이스와 맵은 내용 비교 불가(흠...)
	// 포인터는 주소 비교

	arr1 := [3] int {1,2,3}
	arr2 := [3] int {3,2,1}
	fmt.Println(arr1 == arr2) // 배열 원소가 모두 같은지

	arr3 := [] int {2,3,4,5}
	fmt.Println(reflect.TypeOf(arr2))
	fmt.Println(reflect.TypeOf(arr3))
	fmt.Println(arr3 == nil) // 슬라이스를 nil과 비교, 메모리 할당 여부 확인

	map1 := map[string] int {"hellp": 1, "caca": 2}
	fmt.Println(map1)

	ptr1 := &e
	ptr2 := &e
	fmt.Println(ptr1)
	fmt.Println(ptr1 == ptr2)
	fmt.Println(*ptr2)

	channel := make(chan int)
	go func() {
		channel <- 32 // channel 에 int 32 보냄
	}()
	channelValue := <- channel
	fmt.Println(channelValue)

}
