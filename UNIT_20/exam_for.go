package main

import "fmt"

/**
FizzBuzz 구현하기
1에서 100 까지 출력
3의 배수는 Fizz 출력
5의 배수는 Buzz 출력
3과 5의 공배수는 FizzBuzz 출력

흑 안좋은 코드 같은데...
*/

func main() {
	//EchoFizzBuzz()
	BottleOfBeer()
}

func EchoFizzBuzz() {
	for i := 1; i <= 100; i++ {
		var printed bool = false
		if i%3 == 0 {
			fmt.Print("Fizz")
			printed = true
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
			printed = true
		}

		if printed {
			fmt.Println(" : ", i)
		}
	}
}

func BottleOfBeer() {
	exp := ""
	for i := 99; i >= 0; i-- {
		switch {
		case i > 1:
			exp = "bottles"
		case i == 1:
			exp = "bottle"
		default:
			exp = ""
		}

		if exp == "" {
			fmt.Println("No more bottles of beer on the wall...")
		} else {
			fmt.Println(i, exp, "of beer on the wall")
		}
	}
}
