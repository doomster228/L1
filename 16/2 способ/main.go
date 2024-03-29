package main

import "fmt"

func main() {
	arr := []int{9, 3, 6, 2, 8, 1, 5, 4, 7}
	fmt.Println("Исходный массив:", arr)

	quicksort(arr)
	fmt.Println("Отсортированный массив:", arr)
}

func quicksort(arr []int) {
	// Если длина массива меньше или равна 1, то функция завершает работу
	if len(arr) <= 1 {
		return
	}

	// Создаём переменную, в которую записываем последний элемент массива
	pivot := arr[len(arr)-1]

	// Создаём счётчик для отслеживания индекса, до которого элементы меньше или равные pivot были уже перемещены
	i := 0

	// Просматриваем все элементы массива, кроме последнего и, если элемент меньше или равен pivot, то меняем его местами с элементом i
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Меняем местами элементы массива, чтобы pivot оказался на своем месте
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	// Рекурсивно вызываем эту же функцию для сортировки левой и правой частей массива. Левая часть - это элементы до индекса i, а правая часть - это элементы после индекса i
	quicksort(arr[:i])
	quicksort(arr[i+1:])
}
