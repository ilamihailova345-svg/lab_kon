package main

import (
	"fmt"
	"time"
)

func generate(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// читает из канала через range и выводит, кто какое число забрал
func worker(id int, ch chan int) {
	// range будет читать, пока канал не закроется
	for n := range ch {
		fmt.Printf("Worker %d получил число %d\n", id, n)
	}
}

func main() {
	ch := make(chan int)

	go generate(ch)

	// читают из 1 канала
	// кто первый прочитал — тот и забрал число
	go worker(2, ch)

	// ожидает 1 секунду, чтобы воркеры успели прочитать все числа
	time.Sleep(1 * time.Second)
}