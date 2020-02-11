package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	arr2 := [5]float32{2, 3, 4, 5, 6}
	for i := 0; i < 5; i++ {
		fmt.Println(arr2[i])
	}

	arr3 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(arr3[4])

	arr4 := [...]string{"adad", "dsds", "Ewwr"}
	fmt.Println(arr4[2])

	arr5 := []int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(arr3)) // 상수 배열
	fmt.Println(reflect.TypeOf(arr4)) // 상수 배열
	fmt.Println(reflect.TypeOf(arr5)) // 가변 배열 ...? map?

	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	for key, value := range arr5 {
		fmt.Println(key, value)
	}

	arr5_copy := arr5
	for key, _ := range arr5_copy {
		if arr5[key] != arr5_copy[key] {
			fmt.Println("false!")
		}
	}
}
