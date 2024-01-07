package main

import (
	"fmt"

	"github.com/sudhirj/slicy"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный срез:", slice)

	var index int
	fmt.Print("Введите индекс элемента, который нужно удалить: ")
	fmt.Scan(&index)

	// Проверяем, что индекс находится в пределах среза
	if index >= 0 && index < len(slice) {
		// Удаляем i-й элемент из среза
		newSlice := slicy.Remove[[]int, int](slice, func(value int, i int, slice []int) bool {
			return i == index
		})
		fmt.Println("Срез после удаления элемента:", newSlice)
	} else {
		fmt.Println("Некорректный индекс")
	}

}
