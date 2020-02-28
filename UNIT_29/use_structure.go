package main

import (
	"fmt"
	"reflect"
)

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	type InnerStrc struct {
		hi string
	}

	rect := Rectangle{3, 5}
	rect.Width = 4
	fmt.Println(rect)

	outerFunc()

	innerStrc1 := new(InnerStrc) // 인스턴스 생성하면서 값 초기화는 불가능하다고 하네요
	innerStrc1.hi = "hihi"       // 구조체포인터도 . 으로 접근
	innerStrc2 := InnerStrc{hi: "adad"}
	fmt.Println(innerStrc1)
	fmt.Println(innerStrc2)

	rect2 := getRectangle(3, 5)
	rect3 := getRectangle(7, 8)
	fmt.Println(&rect2)
	fmt.Println(&rect3)
}

func outerFunc() {
	//strc := InnerStrc // 범위 적용되네
	strc := Rectangle{}
	strc.Width = 3
	fmt.Println(strc)
}

// go 언어에서는 지역 변수를 계속 참조하고 있다면 스코프 외부에서도 변수가 해제되지 않는다고 하네요
// 참고로 파라미터나 반환형이 구조체포인터가 아닌 구조체면 복사되니까 신경쓰기
func getRectangle(w int, h int) *Rectangle {
	rect := Rectangle{w, h}
	fmt.Println(reflect.TypeOf(rect))
	return &rect // 지역구조체의 주소를 준다고 이게 동작하네
}
