package main

import "fmt"

func main() {

	//1

	const num = 10

	if num < 0 {
		fmt.Println("Отрицательное")
	} else if num == 0 {
		fmt.Println("Ноль")
	} else if num > 0 {

		if num%2 == 0 {
			fmt.Println("Четное")
		} else if num != 0 {
			fmt.Println("Нечетное")
		}

	}

	//2

	const price = 500
	const wallet = 650

	if change := wallet - price; change >= 0 {
		fmt.Printf("Куплено, ваша сдача - %d\n\n", change)
	} else if change < 0 {
		fmt.Println("Не хватает денег")
	}

	month := 4

	switch month {
	case 12, 1, 2:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Ошибка: неверный месяц")
	}

	signal := "green"
	pedestrians := true

	switch {
	case signal == "green" && pedestrians != true:
		fmt.Println("Можно ехать")
	case signal == "green" && pedestrians != false:
		fmt.Println("Жди, пока пройдут пешеходы")
	case signal == "yellow":
		fmt.Println("Приготовься")
	case signal == "red":
		fmt.Println("Стой!")
	default:
		fmt.Println("Светофор сломался, движение по знакам")
	}
}
