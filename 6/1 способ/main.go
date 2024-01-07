package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаём канал для остановки горутины
	done := make(chan bool)

	// Запускаем горутину
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина остановлена")
				// Остановка выполнения горутины
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(time.Second)
			}
		}
	}()

	// Позволяем горутине работать в течение 5 секунд
	time.Sleep(5 * time.Second)

	// Отправляем сигнал для остановки горутины
	done <- true

	// Ждём пока горутина корректно завершит работу
	time.Sleep(time.Second)

	fmt.Println("Основная горутина завершена")
}
