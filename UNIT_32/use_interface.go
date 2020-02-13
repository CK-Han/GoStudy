package main

import (
	"fmt"
	"strconv"
)

type Hello interface {
}

type MyInt int // typedef 와 유사한 동작인듯
type MyFloat32 float32

func (i MyInt) print() {
	fmt.Println(i)
}

//func (i int) print2() {} 빌트인 자료형은 안된다네
func (f MyFloat32) print() {
	fmt.Printf("%06f\n", f)
}

type Rect struct {
	Width  int
	Height int
}

func (rect *Rect) print() {
	fmt.Println(rect.Width, rect.Height)
}

type Printer interface {
	print() // 이 함수는 3가지 타입에 대해 연결되어있음, 함수의 시그니처는 동일해야함 (파라미터, 리턴 포함)
}

// 덕 타이핑?
// 값이나 인스턴스의 실제 타입에 상관 없이, 구현된 메서드로만 타입을 판단하는 방싟
type Duck struct {
}

func (d *Duck) quack() {
	fmt.Println("꽥!")
}
func (d *Duck) feathers() {
	fmt.Println("오리는 흰색과 회색 털을 가집니다")
}

type People struct {
}

func (p *People) quack() {
	fmt.Println("사람이 꽥!")
}
func (p *People) feathers() {
	fmt.Println("사람이 털?")
}

type Quacker interface {
	quack()
	feathers()
}

// 템플릿 함수 느낌
func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

// 빈 인터페이스로 모든 타입을 받음
func formatString(arg interface{}) string {
	switch arg.(type) { // type 이라는 키워드라... switch 문에서만 가능
	case int:
		i := arg.(int) // 이렇게 타입을 지정해서 값을 가져와야함
		return strconv.Itoa(i)
	case string:
		s := arg.(string)
		return s
	case Rect:
		rect := arg.(Rect)
		return strconv.Itoa(rect.Width) + " " + strconv.Itoa(rect.Height)
	case *Rect:
		rect := arg.(*Rect)
		return strconv.Itoa(rect.Width) + " " + strconv.Itoa(rect.Height)
	default:
		return arg.(string)
	}
}

func main() {
	var h Hello // nil
	_ = h

	var i MyInt = 3
	var f MyFloat32 = 3.1414
	var rect = Rect{3, 5}
	var p Printer
	p = i // i 를 인터페이스 p 에 대입 ...? 인터페이스에 자료형을 담는다는 느낌임
	p.print()
	p = f // MyFloat32 연결
	p.print()
	p = &rect
	p.print()

	pArr := []Printer{i, f, &rect}
	for _, value := range pArr {
		value.print()
	}

	duck := Duck{}
	people := People{}
	inTheForest(&duck)
	inTheForest(&people)

	// 타입이 해당 인터페이스를 구현하는지 검사 interface{}(instance).(interface)
	var donald Duck
	v, ok := interface{}(&donald).(Quacker)
	fmt.Println(v, ok)

	fmt.Println(formatString(1))
	//fmt.Println(formatString(3.55)) 안도네 string 으로 형변환 해줄줄 알았는데 암시적 변환 고딴거 없음
	fmt.Println(formatString("asdasd"))
	fmt.Println(formatString(Rect{3, 5}), formatString(&Rect{7, 10}))

	//인터페이스 변수가 특정 타입인지 검사
	p3 := interface{}(Rect{3, 5})
	v2, ok2 := p3.(Rect)
	fmt.Print(v2, ok2)
}
