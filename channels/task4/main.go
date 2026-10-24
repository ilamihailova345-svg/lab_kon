package main

import "fmt"

// 1. пишет числа 1..5 в первый канал
func generate(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// 2. читает из in, возводит в квадрат и пишет в out
func square(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// 3. читает из канала и выводит на экран
func print(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	chan1 := make(chan int) // generate - square
	chan2 := make(chan int) // square - print

	// первые 2 этапа в горутинах
	go generate(chan1)
	go square(chan1, chan2)

	// 3 этап вызываем в main
	// конвейер синхронизируется без Sleep
	// main будет читать, пока print не закончит, а print закончит,
	// когда chan2 закроется
	print(chan2)
}
