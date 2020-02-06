package main

import "fmt"

func main() {
	const (
		SUNDAY 		= 1
		MONDAY		= 2
		TUESDAY		= 3
	)

	const (
		SUN		= iota // 0
		MON
		TUE
	)
	fmt.Println(TUE)

	const (
		ONE		= 1 << iota
		TWO		= 1 << iota
		THREE	= 1 << iota
		FOUR	= 1 << iota
	)
	fmt.Println(FOUR)

	const (
		bit0, mask0		= 1 << iota, 1 << iota - 1
		bit1, mask1		// 2, 1
		_, _			// 4, 3 이 생략
		bit3, mask3		// 8, 7
	)

	fmt.Println(bit3)
	fmt.Println(mask3)
}
