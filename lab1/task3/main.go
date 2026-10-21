package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		// передаём i как аргумент (id), а не используем замыкание
		go func(id int) {
			fmt.Printf("Goroutine %d says hello\n", id)
		}(i)
	}

	
	// НАБЛЮДЕНИЯ:
	// Программа работает дольше
	// Выводется меньше 10000, потому что 500мс может не хватить 
	// Порядок вывода — случайный

	// Это доказывает, что time.Sleep — не очень надежный способ ожидания
	time.Sleep(500 * time.Millisecond)
}