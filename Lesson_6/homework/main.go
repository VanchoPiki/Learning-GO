package main

import "fmt"

func main() {
	//1

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("Число %d\n", i)
		}
	}

	//2

	i := 0

	for deposit := 10_000.0; deposit <= 20_000.0; i++{
		deposit = deposit * 1.08
	}
	fmt.Println(i)

	//3

	for i := 1; i <= 15; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//4

	num := 1

	for {
		num++
		if num%7 == 0 && num%11 == 0 {
			fmt.Println("Нашел число -", num)
			break
		}
		fmt.Println(num)
	}
}
