package main

import "fmt"

// генерирует числа от 1 до N и отправляет в канал
// Канал out — только для записи, чтобы функция не могла из него читать
func inputStream(n int, out chan<- int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

// читает из in, умножает на 2 и отправляет в out
func doubler(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

// читает из канала и выводит результат
func printer(in <-chan int) {
	for num := range in {
		fmt.Printf("Результат: %d\n", num)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	// 1, 2 этапа в горутинах - они работают параллельно
	go inputStream(10, chan1)
	go doubler(chan1, chan2)

	// 3 этап запускаем в главной горутине
	printer(chan2)


	// В PHP я бы просто создала массив [1..10], прошлась по нему циклом foreach,
	// умножила каждое число на 2 и вывела на экран. Всё работало бы последовательно.
	// В Go каналы сами синхронизируют потоки, а горутины легче потоков.
}