package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Создаём каналы для остановки горутины
	done := make(chan os.Signal, 1)
	stop := make(chan bool)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// Запускаем горутину
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина остановлена")
				// Остановка выполнения горутины
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(time.Second)
			}
		}
	}()

	// Ожидаем Ctrl+C для остановки горутины
	<-done

	// Отправляем сигнал остановки горутины
	stop <- true

	// Ждём пока горутина корректно завершит работу
	time.Sleep(time.Second)

	fmt.Println("Основная горутина завершена")
}
