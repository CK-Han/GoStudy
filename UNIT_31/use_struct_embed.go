package main

import "fmt"

/**
Go 언에서는 클래스가 없으니 상속도 없지만
embedding 을 통해 상속과 같은 효과를 낼 수 있답니다
*/

type People struct {
	name string
}

func (p *People) greeting() {
	fmt.Println("i'm", p.name, "hihi~")
}

type Student struct {
	people People // has-a
	id     int
}

type AnotherStudent struct {
	People // 필드명 없이(익명 필드) 타입만 선언 시 is-a
	id     int
}

// 오버라이딩 가능
func (st *AnotherStudent) greeting() {
	fmt.Println("i'm,", st.name, " another hihi~")
}

func main() {
	student1 := Student{}
	student1.people.greeting()

	student2 := AnotherStudent{People{"aa"}, 3} // 이거보다 좋은 방법 없나..?
	student2.greeting()
}
