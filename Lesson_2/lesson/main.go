package main

import "fmt"

func main() {
	var rating float64 = 8.6
	var title string = "OInt"
	var year int = 2014
	var watched bool = true

	fmt.Println(rating, title, year, watched)

	var rating2 = 8.6
	var title2 = "OInt"
	var year2 = 2014
	var watched2 = true

	fmt.Println(rating2, title2, year2, watched2)

	rating3 := 8.6
	title3 := "OInt"
	year3 := 2014
	watched3 := true

	fmt.Println(rating3, title3, year3, watched3)

	const tit = 15
	const (
		tit1 = 10
		tit2 = 15
		tit3 = 20
	)

	zov, ww, tt := 1488, 12, 14
	fmt.Println(zov, ww, tt)

	left, right := "left", "right"
	fmt.Println(left, right)
	left, right = "right", "left"
	fmt.Println(left, right)
}
