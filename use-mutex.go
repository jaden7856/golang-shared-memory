package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type muCounter struct {
	i  int64
	mu sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
}

func (c *muCounter) increment() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i += 1      // 공유 데이터 변경
	c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금 해제
}

func (c *muCounter) display() {
	fmt.Println(c.i)
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := muCounter{i: 1}
	done := make(chan struct{})

	for i := 0; i < 100; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
		c.display()
	}

	for i := 0; i < 100; i++ {
		<-done
	}

	endTime := time.Since(start)
	fmt.Printf("실행시간 : %s\n", endTime)
}
