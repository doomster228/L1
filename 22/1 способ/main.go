package main

import (
	"fmt"
)

// Этот способ сделан исходя из обозначенного в задании числа 2^20, операции над которым не требуют использования big.Int и big.Float
func main() {
	// Инициализируем и вводим переменные
	var a, b int
	fmt.Print("Введите значение a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение b: ")
	fmt.Scan(&b)

	// Умножение
	multiplication := a * b
	fmt.Printf("Умножение: %d\n", multiplication)

	// Деление
	division := float64(a) / float64(b)
	fmt.Printf("Деление: %g\n", division)

	// Сложение
	addition := a + b
	fmt.Printf("Сложение: %d\n", addition)

	// Вычитание
	subtraction := a - b
	fmt.Printf("Вычитание: %d\n", subtraction)
}
