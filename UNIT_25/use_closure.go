package main

import "fmt"

func main() {
	f1 := func(int) int {
		return 1
	}
	fmt.Println(f1) // 함수 포인터 개념이므로 주소값 출력

	myTest1 := myTest()
	myTest2 := myTest()
	fmt.Println(myTest1())
	fmt.Println(myTest2())
	fmt.Println(myTest2())
	fmt.Println(myTest2())

}

func myTest() func() (int, []int) {
	i := 5
	arr := []int{1, 2, 3}
	return func() (int, []int) {
		i++
		arr[0] += 1
		return i, arr
	}
}
