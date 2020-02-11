package main

import (
	"fmt"
	"reflect"
)

/**
@brief 슬라이스는 배열과는 달리 레퍼런스 타입입니다
*/

func main() {
	// 크기 선언시 모두 0 으로 초기화
	var arr1 []int = make([]int, 5)
	fmt.Println(reflect.TypeOf(arr1))
	//fmt.Println(arr1[6]) // runtime error

	var arr2 = make([]int, 3)
	arr3 := make([]string, 3)
	for _, value := range arr2 {
		fmt.Println(value)
	}
	_, _ = arr2, arr3

	arr4 := []int{
		3,
		5,
	}
	_ = arr4

	arr5 := make([]int, 5, 10) // reserve 가 10 인 슬라이스
	fmt.Println(arr5[4])
	//fmt.Println(arr5[5]) // reserve 이더라도 런타임에러
	fmt.Println(len(arr5), cap(arr5))
	arr5 = append(arr5, 1, 2)
	arr5Append1 := make([]int, 2)
	arr5 = append(arr5, arr5Append1...) // 슬라이스 를 이을시 ... 표기
	fmt.Println(len(arr5))
	arr5Append2 := make([]int, 3)
	arr5 = append(arr5, arr5Append2...) // 9 + 3 이 된다면?
	fmt.Println(len(arr5), cap(arr5))   // 자동으로 할당했다

	// 배열과 슬라이스의 차이 - 레퍼런스 타입
	arr6 := [3]int{1, 2, 3}
	var arr7 [3]int
	arr7 = arr6
	arr7[0] = 9
	fmt.Println(arr6, arr7)

	arr8 := []int{1, 2, 3}
	arr9 := []int{4, 5, 6}
	arr9 = arr8 // 대입 시 포인터값(참조) 만 복사
	arr9[0] = 9
	fmt.Println(arr8, arr9)

	arr10 := []int{1, 2, 3}
	arr11 := []int{4, 5, 6}
	copy(arr11, arr10) // 값 복사는 copy 함수
	fmt.Println(&arr10[0], &arr11[0])
	arr11[0] = 9
	fmt.Println(arr10, arr11)

	arr12 := []int{1, 2, 3, 4, 5}
	arr13 := arr12[0:2] // [ ... ) 참고로 레퍼런스 복사임
	fmt.Println(arr13)
	arr13[0] = 9
	fmt.Println(arr12)    // 레퍼런스 복사됨
	fmt.Println(arr12[:]) // 시작과 끝 생략 가능

	arr14 := [5]int{1, 2, 3, 4, 5} // 슬라이스가 아닌 배열
	arr15 := arr14[:3]             // 배열로 부터 '슬라이스' 를 만드는 것
	arr15[0] = 9
	fmt.Println(arr14) // 즉, 배열로부터 슬라이스 만들면 레퍼런스 참조된다

	fmt.Println(len(arr15), cap(arr15))
	arr16 := arr15[:2:4] // 3번째 표기로 reserve 공간 설정
	//arr16 := arr15[0:2:15] // 기존 reserve 보다 크면 런타임 에러
	fmt.Println(len(arr16), cap(arr16))

}
