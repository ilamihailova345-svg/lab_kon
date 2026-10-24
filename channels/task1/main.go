package main

import "fmt"

func main() {
	// небуферизированный канал для строк (отправитель блокируется, пока получатель не прочитает)
	ch := make(chan string)

	// запускаем анонимную горутину
	go func() {
		// горутина ЗАБЛОКИРУЕТСЯ здесь, пока main не начнёт читать
		ch <- "Hello, World!"
	}()

	// читаем из канала и выводим
	message := <-ch
	fmt.Println(message)
}