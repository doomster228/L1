package main

import (
	"fmt"
)

// Вычисление квадратов чисел с использованием каналов
func calculateSquare(numbers []int, results chan<- int) {
	// Вычисляем квадрат каждого числа из массива и отправляем результат и исходное число в канал
	for _, num := range numbers {
		square := num * num
		results <- square
		results <- num
	}

	// Закрываем канал
	close(results)
}

func main() {
	// Создаём требуемый массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов вычислений между горутинами
	results := make(chan int)

	// Вызываем функцию вычисления квадратов чисел в отдельной горутине
	go calculateSquare(numbers, results)

	// Выводим в консоль кождое из значений, записанных в канал
	for result := range results {
		fmt.Printf("Квадрат числа %d равен: %d\n", <-results, result)
	}
}
