package main

import "fmt"

func main() {

	//1

	// Создаем пустую карту, где ключ - строка, а значение - число
	ages := make(map[string]int)
	fmt.Println(ages)

	prices := map[string]float64{
		"Хлеб":   45.5,
		"Молоко": 89.9,
		"Кофе":   250.0, // Запятая в конце обязательна!
	}
	fmt.Println(prices)

	var myMap map[string]int

	// Читать из нее можно:
	fmt.Println(myMap["test"]) // Выведет 0

	// НО ПИСАТЬ В НЕЕ НЕЛЬЗЯ:
	// myMap["test"] = 10 // ПАНИКА! (panic: assignment to entry in nil map)

	//2
	users := make(map[string]string)

	// 1. Добавление и Перезапись (синтаксис одинаковый)
	users["admin"] = "Алексей"
	users["guest"] = "Иван"
	users["guest"] = "Дмитрий" // Перезаписали значение для ключа "guest"

	// 2. Чтение
	fmt.Println("Админ:", users["admin"]) // Алексей

	// 3. Чтение НЕСУЩЕСТВУЮЩЕГО ключа
	// Go НЕ падает с ошибкой, а возвращает НУЛЕВОЕ значение типа (для string это "")
	fmt.Println("Модератор:", users["moderator"]) // Выведет пустоту ""

	// 4. Удаление через функцию delete(карта, ключ)
	delete(users, "guest")
	fmt.Println("После удаления:", users) // map[admin:Алексей]

	// Удаление несуществующего ключа ошибки НЕ вызывает:
	delete(users, "some_ghost") // Безопасно, ничего не произойдет

	//3

	scores := map[string]int{
		"Иван": 0, // Иван есть, но набрал 0
	}

	// Карта при чтении возвращает ДВА значения: (значение, был_ли_ключ)
	score, ok := scores["Иван"]
	fmt.Printf("Очки: %d, Найден: %t\n", score, ok) // Очки: 0, Найден: true

	score, ok = scores["Петр"]
	fmt.Printf("Очки: %d, Найден: %t\n", score, ok) // Очки: 0, Найден: false

	// Красивая идиоматичная проверка через if с инициализацией:
	if val, ok := scores["Петр"]; ok {
		fmt.Println("Петр набрал:", val)
	} else {
		fmt.Println("Студент Петр не найден!")
	}

	//4

	fruits := map[string]int{
		"Яблоки": 10,
		"Груши":  5,
		"Сливы":  20,
	}

	for key, value := range fruits {
		fmt.Printf("Фрукт: %s, Количество: %d\n", key, value)
	}
}
