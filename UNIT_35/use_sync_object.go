package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//appendFunc()
	//appendWithMutex()
	//useRWMutex()
	//useCondition()
	//useOnce()
	//usePool()
	//useWaitGroup()
	useAtomic()
}

func appendFunc() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)

	arr := []int{}

	go func() {
		for i := 0; i < 1000; i++ {
			arr = append(arr, 1)
			runtime.Gosched() // cpu 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			arr = append(arr, 1)
			runtime.Gosched() // cpu 양보
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println(len(arr))
}

func appendWithMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	arr := []int{}
	mutex := new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			arr = append(arr, 1)
			mutex.Unlock()
			runtime.Gosched() // cpu 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			arr = append(arr, 1)
			mutex.Unlock()
			runtime.Gosched() // cpu 양보
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println(len(arr))
}

func useRWMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := 1
	rwMutex := new(sync.RWMutex)

	// read lock(RLock)은 읽는 동작은 막지 않음, 대신 읽는 중 값이 바뀌면 안되니 쓰기 락이 걸림
	// write lock 은 쓰는 중 다른곳에서 읽거나 쓰면 안되므로 기존 mutex(Lock) 과 동일

	go func() { // write goroutine
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() { // read goroutine
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read :", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() { // read goroutine
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read :", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(time.Second * 10)
}

func useCondition() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("wait begin :", n)
			cond.Wait() // 조건변수 대기, 이 때 내부에서 등록했던 잠금객체를 unlock 하고 대기하네
			fmt.Println("wait end :", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 3개가 모두 wait 상태로 갈때까지 대기
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
	}

	//mutex.Lock()
	//cond.Broadcast()
	//mutex.Unlock()

	fmt.Scanln()
}

func useOnce() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(func() {
				fmt.Println("hihi~")
			})
		}(i)
	}

	fmt.Scanln()
}

func usePool() {
	type Data struct {
		tag    string
		buffer []int
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	pool := sync.Pool{
		New: func() interface{} { // Get() 호출시 사용되는 함수, 풀에 객체가 없을 시
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)
			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}
			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()
}

func useWaitGroup() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	for i := 1; i <= 10; i++ {
		wg.Add(2) // 그룹 개수 추가
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)

		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i * 2)
	}

	wg.Wait()
	fmt.Println("end!")

	wg2 := new(sync.WaitGroup)
	for i := 100; i < 110; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			fmt.Println(n)
		}(i)
	}
	wg2.Wait()
}

func useAtomic() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
