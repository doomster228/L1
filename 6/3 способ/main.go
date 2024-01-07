package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем новый контекст и функцию для его отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину
	go func() {
		for {
			select {
			default:
				fmt.Println("Горутина работает")
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println("Горутина остановлена")
				// Остановка выполнения горутины
				return
			}
		}
	}()

	// Позволяем горутине работать в течение 5 секунд
	time.Sleep(5 * time.Second)

	// Останавливаем выполнение горутины
	cancel()

	// Ждём пока горутина корректно завершит работу
	time.Sleep(time.Second)

	fmt.Println("Основная горутина завершена")
}
