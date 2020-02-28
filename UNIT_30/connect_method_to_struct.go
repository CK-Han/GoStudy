package main

import "fmt"

/**
Go 언어에는 클래스가 없는 대신 구조체에 메서드를 연결한답니다!
*/

// 메소드가 따로 없는 Rect 구조체
type Rect struct {
	Width  int
	Height int
}

func main() {
	rect := Rect{3, 5}
	fmt.Println(rect.getArea()) // 메소드가 생겼다

	fmt.Println(rect.getAreaResult())
}

func (rect *Rect) getAreaResult() int {
	//return rect.getAreaResult() // 따로 컴파일 오류 없음
	return rect.getArea() // 함수 선언 위치가 상관 없으니 메소드도 마찬가지구나
}

// 리시버명 *구조체
func (rect *Rect) getArea() int {
	return rect.Width * rect.Height
}

// 인스턴스에 대해 동작 vs 변수에 대해 동작
// static func 같은 느낌은 어떻게 낼지 궁금하네...
func (rect Rect) scale(factor int) {
	rect.Height *= factor
	rect.Width *= factor
}

// 요런식으로 리시버 변수를 사용 안하게 하나?
func (_ Rect) dummy() {
	fmt.Println("dummy")
}
