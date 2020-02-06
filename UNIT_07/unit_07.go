package main

import _ "fmt" // 사용하지 않는 패키지 컴파일 에러 방지

func main() {
	var i int
	i = 3
	var s = "cacaca"
	var k = 35.5
	j := 33

	var idx, name = 3, "aa"
	shortidx, sname := 5, "dd"

	// 사용하지 않은 변수 및 패키지 지우기
	_ = i
	_ = s
	_ = k
	_ = j
	_, _, _, _ = idx, name, shortidx, sname
}
