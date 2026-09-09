package main

import "fmt"

func main() {
	age := 18

	if age < 14 {
		fmt.Println("Ты еще ребенок")
	} else if age < 18 {
		fmt.Println("Ты подросток")
	} else {
		fmt.Println("Ты взрослый!")
	}

	if length := len("Hello, Go!"); length > 5 {
		fmt.Printf("Строка длинная, в ней %d символов\n", length)
	} else {
		fmt.Printf("Строка короткая, всего %d символов\n", length)
	}

	day := "Суббота"

	switch day {
	case "Понедельник", "Вторник", "Среда", "Четверг", "Пятница":
		fmt.Println("Будний день, пора работать.")
	case "Суббота", "Воскресенье":
		fmt.Println("Выходной! Можно отдыхать.")
	default:
		fmt.Println("Неизвестный день недели...")
	}

	score := 85

	switch { // Условия нет!
	case score >= 90:
		fmt.Println("Отлично: A")
	case score >= 75:
		fmt.Println("Хорошо: B")
	case score >= 60:
		fmt.Println("Удовлетворительно: C")
	default:
		fmt.Println("Не сдал: F")
	}

	number := 1
	switch number {
	case 1:
		fmt.Println("Один")
		fallthrough // Принудительно передает управление в следующий case
	case 2:
		fmt.Println("Два (сработало из-за fallthrough!)")
	}

}
