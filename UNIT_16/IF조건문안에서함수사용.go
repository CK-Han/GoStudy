package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// GO 가 바라보는 경로 개념을 잘 모르겠네...
	b, err := ioutil.ReadFile("./UNIT_16/hello.txt")
	if err == nil {
		fmt.Printf("%s", b)
	}

	// 스코프가 if 문에 묶임. else if 및 else 에서도 사용 가능
	if c, er := ioutil.ReadFile("./UNIT_16/hello.txt"); er == nil {
		fmt.Printf("%s", c)
	}
	fmt.Println(b)
	//fmt.Println(c)
}
