package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Указываем количество воркеров
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	// Создаем канал для передачи данных
	dataChannel := make(chan string)

	// Создаем канал для связи между главной программой и обработчиком сигналов
	done := make(chan os.Signal, 1)

	// Создаем обработчик сигналов, которые могут быть использованы для прерывания программы
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем воркеры
	for i := 0; i < workerCount; i++ {
		go worker(i, dataChannel)
	}

	for {
		// Слушаем сигналы операционной системы, чтобы завершить программу по Ctrl+C
		select {
		case <-done:
			// Программа завершает работу по сигналу Ctrl+C
			fmt.Println("Программа завершена по сигналу Ctrl+C")
			return
		default:
			// Ввод строки вручную
			var data string
			fmt.Scanln(&data)
			dataChannel <- data

			// Запись в канал готовой строки
			//dataChannel <- "Hello, world!"
		}
	}
}

// Воркер для обработки данных
func worker(id int, dataChannel <-chan string) {
	for data := range dataChannel {
		// Воркер получает и выводит данные
		fmt.Printf("Воркер %d получил данные: %s\n", id, data)
	}
}
