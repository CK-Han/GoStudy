package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while 대신 for 씀
	j := 5
	for j > 0 {
		fmt.Println(j)
		j--
	}

	_ = j
	CACA:
		//fmt.Println("aaa") // 레이블과 for 사이에 다른 코드 불가
	for i := 0; i < 3; i++ {
		//PAPA:
		for j := 0; j < 3; j++ {
			fmt.Printf("%d %d\n", i, j)
			break CACA // ㄷㄷ; goto 급인데
		}
	}

	// 한 레이블 당 한 for 문
	AAA:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d %d\n", i, j)
			continue AAA
			//break 나 continue 에 레이블을 같이 쓰면
			//해당 레이블에서 break or continue 를 한 것으로 동작
		}
	}

	for i, j := 0, 0; i < 3; i, j = i+1, j+2 {

	}
	// 이러한 변화식은 불가, 위 처럼 대입 나열식으로 가능
	//for i, j := 0, 0; i < 3; i, j = i++, j+2 {
}
