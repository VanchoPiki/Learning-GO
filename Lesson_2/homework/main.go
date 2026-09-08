package main

import "fmt"

func main() {

	const (
		price = 1000.00
		title = "ZOV+"
	)

	var months int = 12
	var discount = 15.0
	total := float64(price) - float64(price)*(discount/100)

	fmt.Println("===== ЧЕК GoFlix =====")
	fmt.Printf("Тариф:%s \nМесяцев:%d \nБазовая цена:%.2f/мес. \nСкидка:%.0f%% \nИтого:%.2f руб./мес.\n", title, months, price, discount, total)
}
