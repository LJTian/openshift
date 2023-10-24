package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

type SpinLock struct {
	flag int32
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.flag, 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.flag, 0)
}

func main() {
	var wg sync.WaitGroup
	var counter int

	strNum := os.Getenv("NUM")

	num, _ := strconv.Atoi(strNum)
	if num == 0 {
		num = 10
		fmt.Println("env NUM is null ,default value 10.")
	} else {
		fmt.Println("env NUM is" + strNum)
	}
	num = num * 1000000

	spinLock := &SpinLock{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			spinLock.Lock()
			counter++
			func() {
				for i := 0; i < 10000; i++ {
					continue
				}
			}()
			spinLock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
