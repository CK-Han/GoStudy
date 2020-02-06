package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var a = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)

	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)


	var com1 complex64 = complex(3, 32323)
	fmt.Println(com1)
	fmt.Println(unsafe.Sizeof(com1))

	var num1 uint8 = 255
	var num2 uint8 = num1 + 32
	fmt.Println(num2)

	/*
	var num1 complex64 = 1 + 2i
	var num2 = 4.2324 + 2.436546i

	fmt.Println(real(num1))
	fmt.Println(imag(num2))

	var str = '한' // 작은 따옴표
	_ = str
	var n1 float32 = 3.2
	var n2 = 2
	n3 := n1 + float32(n2)
	_ = n3

	maxUint32 := math.MaxUint32
	_ = maxUint32

	// 오버플로우 및 언더플로우는 컴파일 에러

	tempNum := 0.005
	fmt.Println(unsafe.Sizeof(tempNum))
	 */
}
