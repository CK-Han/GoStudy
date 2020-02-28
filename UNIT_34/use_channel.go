package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//channelStart()
	//bufferChannel()
	//rangeAndCloseChannel()
	//forwardChannel()
	selectChannel()
}

func channelStart() {
	done := make(chan bool)
	count := 3
	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("id:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	func() {
		for i := 0; i < count; i++ {
			<-done
			fmt.Println("func:", i)
		}
	}()
}

func bufferChannel() {
	runtime.GOMAXPROCS(1)

	// 동기 채널은 보내는 쪽은 받을때까지 대기, 받는 쪽은 값이 들어올때까지 대기
	// 버퍼 size 설정하면 비동기 채널 생성
	// 보내는 쪽은 버퍼가 가득 찬다면 실행을 멈춤
	// 받는 쪽은 기존처럼 버퍼에 데이터 있을때까지 대기
	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("goroutine called:", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("func:", i)
	}

	fmt.Scanln()
}

func rangeAndCloseChannel() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			if i > 2 {
				time.Sleep(time.Second * 1)
			}

			c <- i
		}
		close(c)
	}()

	// channel range 는 해당 채널이 close 될때까지 순환
	for i := range c {
		fmt.Println(i)
	}

	num, ok := <-c // close 된 채널은 [default 값, false]
	fmt.Println(num, ok)

	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	c <- 3 // close 된 채널에 값 보낼 시 패닉 발생
}

func forwardChannel() {
	c := make(chan int)

	go func(c chan<- int) { // 보내기 전용
		for i := 0; i < 5; i++ {
			c <- i
		}

		c <- 100
		//fmt.Println(<-c) // send_only channel 은 꺼내려할 시 컴파일 에러
	}(c)

	go func(c <-chan int) { // 받기 전용
		for i := range c {
			fmt.Println(i)
		}
		//c <- 1 // receive_only channel 은 보내려할 시 컴파일 에러
	}(c)

	time.Sleep(time.Second * 3)

	c2 := func(n1 int, n2 int) <-chan int {
		recvChan := make(chan int)
		go func() {
			recvChan <- n1 + n2
		}()

		return recvChan
	}(3, 5)

	fmt.Println(<-c2)
}

func selectChannel() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "hihi"
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	//c3 := time.After(time.Millisecond * 180)

	go func() {
		for {
			select {
			//case <-c3:
			//	fmt.Println("time.After called")
			case i := <-c1:
				fmt.Println("c1:", i)
			case s := <-c2:
				fmt.Println("c2:", s)
			}
		}
	}()

	fmt.Scanln()
}
