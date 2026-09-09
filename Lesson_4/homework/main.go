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
}
