package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s4 string = "\U0000d55c\U0000ae00"
	fmt.Println(s4)

	var s5 string = "안녕하123셍"
	fmt.Println(utf8.RuneCountInString(s5))

	var s6 string = "한글"
	fmt.Println(s4 == s6)
	fmt.Println(s4 + "    " + s5)
	fmt.Println(string(s5[10]))
	for idx, ch := range s5 {
		_= idx
		fmt.Println(string(ch))
	}

	//s5[10] = 'k'  컴파일 에러(문자열 직접 수정 불가)

	var s7 = "\ud55c\uae00\n"
	fmt.Println(s7)
	fmt.Println(utf8.RuneCountInString(s7))
}