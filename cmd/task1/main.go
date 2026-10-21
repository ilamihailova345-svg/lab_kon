package main

import (
	"fmt"
	"time"
)

// печатает числа от 1 до n с паузами 100мс
func printNumbers(id int, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start lab_kon")
	printNumbers(1, 5)
	fmt.Println("End lab_kon")
}
