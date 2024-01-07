package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Указываем количество воркеров
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	// Создаем канал для передачи данных и ожидающую группу
	dataChannel := make(chan string)
	var wg sync.WaitGroup

	// Создаем новый контекст и функцию для его отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Делаем количество элементов в группе равным количеству воркеров
	wg.Add(workerCount)

	// Запускаем воркеры
	for i := 0; i < workerCount; i++ {
		go worker(dataChannel, &wg, ctx, i)
	}

	// Создаем канал для связи между главной программой и обработчиком сигналов
	done := make(chan os.Signal, 1)

	// Создаем обработчик сигналов, которые могут быть использованы для прерывания программы
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Слушаем сигналы операционной системы, чтобы завершить программу по Ctrl+C
	go func() {
		<-done
		// Программа завершает работу по сигналу Ctrl+C
		fmt.Println("Программа завершена по сигналу Ctrl+C")
		cancel()
	}()

	for {
		// Ввод строки вручную
		var data string
		fmt.Scanln(&data)
		select {
		case <-ctx.Done():
			// Программа завершает работу по сигналу Ctrl+C
			wg.Wait()
			return
		default:
			dataChannel <- data
			// Запись в канал готовой строки
			//dataChannel <- "Hello, world!"
		}
	}
}

func worker(dataChannel chan string, wg *sync.WaitGroup, ctx context.Context, id int) {
	defer wg.Done()
	for {
		select {
		case data := <-dataChannel:
			// Выводим данные в stdout
			fmt.Printf("Воркер %d получил данные: %s\n", id, data)
		case <-ctx.Done():
			// Программа завершает работу по сигналу Ctrl+C
			return
		}
	}
}
