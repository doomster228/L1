package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Структура, содержащая отображение для хранения данных и мьютекс для обеспечения безопасности доступа к отображению из разных горутин
type DataStore struct {
	data map[string]string
	mux  sync.Mutex
}

// Устанавливаем значения элементов отображения
func (ds *DataStore) Set(key, value string) {
	ds.mux.Lock()
	defer ds.mux.Unlock()

	ds.data[key] = value
}

// Получаем значения элементов отображения
func (ds *DataStore) Get(key string) (string, bool) {
	ds.mux.Lock()
	defer ds.mux.Unlock()

	value, ok := ds.data[key]
	return value, ok
}

func main() {
	// Создаём экземпляр отображения
	dataStore := &DataStore{
		data: make(map[string]string),
	}

	// Создаём группу ожидания
	var wg sync.WaitGroup

	// Создаём массив ключей
	keys := []string{"key1", "key2", "key3"}

	// Запускаем несколько горутин для записи значений в отображение
	for i, key := range keys {
		wg.Add(1)
		go func(k string, index int) {
			defer wg.Done()
			value := "value" + strconv.Itoa(index+1)
			dataStore.Set(k, value)
		}(key, i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Получаем значения и выводим их на экран
	for _, key := range keys {
		value, ok := dataStore.Get(key)
		if ok {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		} else {
			fmt.Printf("Key: %s not found\n", key)
		}
	}
}
