package main

import f "fmt" // 패키지 별칭
import . "runtime" // 전역 패키지로 설정

func main() {
	f.Println(NumCPU())
}
