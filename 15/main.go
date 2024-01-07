package main

var justString string

func someFunc() {
	// Создаём строку по всей видимости размером 1 МБ
	v := createHugeString(1 << 10)

	// Если длина строки больше 100 символов, то сохраняем в переменную только первые 100 символов, если меньше - всю строку
	if len(v) > 100 {
		justString = v[:100]
	} else {
		justString = v
	}

	// Освобождаем память, занимаемую большой строкой после использования
	v = ""
}

func main() {
	someFunc()
}
