package main

import "fmt"

func main() {
	var ptr1 *int = new(int)
	fmt.Println(ptr1)

	num1 := 3
	ptr2 := &num1
	fmt.Println(&num1, ptr2)
	//ptr2++ go 는 포인터 대입이나 포인터 연산 불가함

	num2 := 15
	callByRef(&num2)
	fmt.Println(num2)
}

func callByRef(num *int) {
	*num++
}
