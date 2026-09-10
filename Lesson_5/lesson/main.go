package main

import "fmt"

// Функция принимает два параметра и ничего не возвращает
// Запись (name string, age int)
func printUserInfo(name string, age int) {
	fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
}

// Фишка: сокращенная запись (a, b int) вместо (a int, b int)
func printSum(a, b int) {
	fmt.Println("Сумма:", a+b)
}

// Функция принимает два float64 и возвращает одно число float64
func multiply(x, y float64) float64 {
	return x * y // Возвращаем результат умножения
}

// Функция возвращает ДВА значения: результат деления и остаток
func divide(a, b int) (int, int) {
	quotient := a / b          // Целая часть
	remainder := a % b         // Остаток
	return quotient, remainder // Возвращаем оба через запятую!
}

func getNameAndAge() (string, int) {
	return "Иван", 30
}

func main() {
	// Вызов функций:
	printUserInfo("Алексей", 25)
	printSum(10, 20)

	// Сохраняем результат работы функции в переменную
	result := multiply(2.5, 4.0)
	fmt.Println("Результат:", result) // 10

	// Принимаем оба значения в две разные переменные
	q, r := divide(10, 3)
	fmt.Printf("Частное: %d, Остаток: %d\n", q, r) // Частное: 3, Остаток: 1

	// Нам нужно только имя, а возраст не интересен:
	name, _ := getNameAndAge()

	fmt.Println("Привет,", name)
	// Переменной для возраста нет, ошибки компиляции тоже нет!
}
