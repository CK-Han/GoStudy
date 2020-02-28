package main

import "fmt"

func main() {
	foo()

	fmt.Println("after panic")
}

func f() {
	// recover는 반드시 defer 함수로 사용해야 함
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	panic([]int{1, 2, 3})
}

func foo() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	arr := []int{1, 2, 3}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
