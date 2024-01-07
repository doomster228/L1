package main

import "fmt"

// Исходная структура
type Human struct {
	name string
	age  int
}

// Метод, привязанный к структуре Human
func (h *Human) introduce() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.name, h.age)
}

// Та структура, в которую вложена исходная
type Action struct {
	Human
}

// Метод, привязанный к структуре Action, в котором задействуется поле из структуры Human
func (a *Action) work() {
	fmt.Printf("%s занимается работой.\n", a.name)
}

func main() {
	// Создаём объекты с помощью каждой из структур
	human := Human{name: "Максим", age: 26}
	action := Action{Human: human}

	// С помощью экземпляра дочерней структуры вызываем методы, привязанные как к дочерней так и к родительской структуре
	action.introduce()
	action.work()
}
