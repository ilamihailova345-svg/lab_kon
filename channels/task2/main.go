package main

import (
	"fmt"
	"math/rand/v2"
)

// отправляет 10 случайных чисел в канал
func generate(ch chan int) {
	for i := 0; i < 10; i++ {
		number := rand.IntN(100)
		ch <- number
	}
	// закрываем канал после отправки всех чисел
	// без close(ch) цикл for rang будет ждать вечно
	// и программа упадёт с ошибкой deadlock
	close(ch)
}

func main() {
	ch := make(chan int)

	// запускаем генератор в отдельной горутине
	go generate(ch)

	//по каналу читает данные, пока канал не будет ЗАКРЫТ
	for n := range ch {
		fmt.Println(n)
	}
}