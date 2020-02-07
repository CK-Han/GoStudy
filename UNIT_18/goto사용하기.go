package main

import "fmt"

func main() {
	for i := 32; i > 0; i /= 2 {
		if i < 0 {
			fmt.Println(i)
			goto RETURN
		}
	}

	RETURN:
		fmt.Println("function returned!")
}
