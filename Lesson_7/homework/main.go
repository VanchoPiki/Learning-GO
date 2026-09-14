package main

import "fmt"

func main() {
	//1

	cities := [4]string{"Москва", "Париж", "Токио", "Берлин"}

	fmt.Println(cities)
	fmt.Println(len(cities))

	cities[0] = "Ростов"

	fmt.Println(cities)
	fmt.Println(cities[0], cities[3])

	//2

	prices := [5]float64{120.5, 45.0, 310.2, 89.9, 15.4}
	price_rang := 0.0
	for _, price := range prices {
		price_rang += price
	}
	sered_price := price_rang / float64(len(prices))
	fmt.Println(prices)
	fmt.Printf("Сумма - %.2f, Средняя цена - %.2f\n", price_rang, sered_price)

	//3

	numbers := [7]int{23, -5, 42, 18, 99, 3, 77}
	maxNum := numbers[0]
	maxInd := 0
	for i, num := range numbers {
		if num > maxNum {
			maxNum = num
			maxInd = i
		}
	}
	fmt.Printf("Максимальное число: %d на позиции %d\n", maxNum, maxInd)

	//4

	myArr := [3]int{1, 2, 3}
	changeFirst(myArr)
	fmt.Println(myArr)
}

// 4
func changeFirst(arr [3]int) {
	arr[0] = 999
}
