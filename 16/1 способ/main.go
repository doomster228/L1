package main

import (
	"fmt"
	"sort"
)

// Структура, содержащая имя и возраст человека
type Person struct {
	Name string
	Age  int
}

func main() {
	// Создаём срез, который будет содержать данные людей
	people := []Person{
		{Name: "Maxim", Age: 26},
		{Name: "Namjoon", Age: 29},
		{Name: "Seokjin", Age: 31},
		{Name: "Yoongi", Age: 30},
		{Name: "Hoseok", Age: 29},
		{Name: "Jimin", Age: 28},
		{Name: "Taehyung", Age: 28},
		{Name: "Jungkook", Age: 26},
	}

	fmt.Println("Исходный список людей:")
	printPeople(people)

	// Сортировка списка людей по возрасту в порядке возрастания
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("\nСписок людей после сортировки по возрасту:")
	printPeople(people)

	// Сортировка списка людей по имени в порядке убывания
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})

	fmt.Println("\nСписок людей после сортировки по имени:")
	printPeople(people)
}

// Выводит в консоль список людей в его текущем состоянии
func printPeople(people []Person) {
	for _, person := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
	}
}
