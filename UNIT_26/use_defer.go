package main

import (
	"fmt"
	"os"
)

// finally 와 유사한 느낌
func main() {
	defer world()
	hello()
	hello()

	deferFor()

	//safeFileRead()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func deferFor() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func safeFileRead() {
	dir, _ := os.Getwd()
	f, err := os.Open(dir + "\\test.log")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	buf := make([]byte, 100)
	if _, err = f.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
