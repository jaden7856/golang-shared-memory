package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	counter := 0
	for value := range ch {
		counter += value
		fmt.Println(counter)
	}

}

func consumer(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- 1
	}
}

func main() {
	start := time.Now()
	ch := make(chan int)

	go producer(ch)
	consumer(ch)

	endTime := time.Since(start)
	fmt.Println("실행시간 : %s", endTime)
}
