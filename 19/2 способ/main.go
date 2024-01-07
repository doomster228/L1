package main

import (
	"bufio"
	"fmt"
	"os"

	"4d63.com/strrev"
)

func main() {
	// Задаём строку, которую будем переворачивать
	var input string
	fmt.Print("Введите строку: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	input = in.Text()

	// Переворачиваем строку с помощью функции из сторонней библиотеки
	fmt.Println("Перевернутая строка:", strrev.Reverse(input))
}
