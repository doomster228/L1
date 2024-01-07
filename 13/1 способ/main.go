package main

import "fmt"

func main() {
	// Инициализируем и вводим переменные
	var a, b int
	fmt.Print("Введите значение a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение b: ")
	fmt.Scan(&b)
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Меняем значения местами
	a, b = b, a

	// Выводим результат
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
