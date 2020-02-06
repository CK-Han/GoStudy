package main

import "fmt"
import "math"
import "unsafe"

func main() {
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)

	const epsilon = 1e-14
	fmt.Println(math.Abs(a - 9.0) <= epsilon)

	var num1 complex64 = 1 + 2i
	var num2 complex128 = 4.2324 + 2.436546i

	fmt.Println(real(num1))
	fmt.Println(imag(num2))

	var str rune = '한' // 작은 따옴표
	_ = str
	var n1 float32 = 3.2
	var n2 int = 2
	n3 := n1 + float32(n2)
	_ = n3

	maxUint32 := math.MaxUint32
	_ = maxUint32

	// 오버플로우 및 언더플로우는 컴파일 에러

	tempNum := 0.005
	fmt.Println(unsafe.Sizeof(tempNum))
}
