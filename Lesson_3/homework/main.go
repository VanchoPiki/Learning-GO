package main

import "fmt"

func main() {
	var time_of_filme = 8521

	var hours = time_of_filme / 3600

	var minuts = time_of_filme % 3600 / 60

	var seconds = time_of_filme % 60

	fmt.Printf("Общее-%d Часы-%d Минуты-%d Секунды-%d\n\n", time_of_filme, hours, minuts, seconds)

	var total = 4_831_838_208

	var gig = float64(total) / 1073741824.00

	fmt.Printf("Всего - %d, Гиг. - %.2f\n", total, gig)
}
