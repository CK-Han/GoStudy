package main

import (
	"fmt"
	"reflect"
)

func main() {
	Hello()

	fmt.Println(sum(3, 5))
	fmt.Println(getSumAndDiff(3, 5))
	// 가변인자는 슬라이스 타입이다!
	fmt.Println(variableSum(2, 3, 5, 34, 6, 437, 34, 38))
	fmt.Println(variableSum([]int{1, 2, 3, 4, 5}...))
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(variableSum(arr1[:2]...))
	fmt.Println(factorial(5))

	myFunc := func(n int) {
		fmt.Println("it`s", n)
	}
	myFunc(3)
	mySumFunc := sum
	fmt.Println(mySumFunc(3, 5))
	myFunc2 := []func(int, int) int{sum, sum2}
	fmt.Println(myFunc2[1](5, 12))

	myFunc3 := map[string]func(int, int) int{
		"sum":  sum,
		"sum2": sum2,
	}
	fmt.Println(myFunc3["sum2"](3, 8))

	func(name string) {
		fmt.Println(name, "hihi!")
	}("cacaru")
}

// 생각해보니 선언이 필요하거나 함수가 메인 위에있거나 할 필요가 없다
func Hello() {
	fmt.Println("hello!")
}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func sum2(n1 int, n2 int) (ret int) {
	ret = n1 + n2
	return
}

func getSumAndDiff(n1 int, n2 int) (sum int, diff int) {
	return n1 + n2, n1 - n2
}

func variableSum(n ...int) (ret int) {
	fmt.Println(reflect.TypeOf(n))
	ret = 0
	for _, i := range n {
		ret += i
	}
	return
}

func factorial(n int64) int64 {
	if n <= 1 {
		return n
	}

	return n * factorial(n-1)
}
