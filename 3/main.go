package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходная последовательность чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем каналы для передачи данных
	input := make(chan int)
	output := make(chan int)

	// Создаем группу горутин
	var wg sync.WaitGroup

	// Запускаем горутину для чтения чисел из входного канала
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	// Запускаем горутины для обработки чисел
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var sum int
			for num := range input {
				sum += num * num
			}
			output <- sum
		}()
	}

	// Запускаем горутину для суммирования результатов
	wg.Add(1)
	go func() {
		defer wg.Done()
		var totalSum int
		for i := 0; i < len(numbers); i++ {
			totalSum += <-output
		}
		fmt.Printf("Сумма квадратов чисел: %d\n", totalSum)
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}
