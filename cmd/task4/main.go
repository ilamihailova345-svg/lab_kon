package main

import (
	"fmt"
	"time"
)

func printNumbers(id int, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start lab1")

	go printNumbers(1, 5)
	go printNumbers(2, 5)
	go printNumbers(3, 5)
	go printNumbers(4, 5)
	go printNumbers(5, 5)


	// используем time.Sleep, чтобы "подождать" горутины
	// это ненадёжно, потому что не знаем, сколько времени нужно
	// не можем гарантировать, что все 5 строк будут напечатаны
	time.Sleep(100 * time.Millisecond)

	fmt.Println("End lab1")
}