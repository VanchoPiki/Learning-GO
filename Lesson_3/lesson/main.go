package main

import "fmt"

func main() {
	var zeroint int
	views := 1_500_000

	fmt.Println(zeroint)
	fmt.Println(views)

	var zerofloat float64
	views2 := 85.6

	fmt.Println(zerofloat)
	fmt.Println(views2)

	var zerostr string
	movtitle := "Title"

	fmt.Println(zerostr)
	fmt.Println(movtitle)

	synk := `dfsdfsdf
dsfsdfsfdsdf
sdfsdfsdfsdf
sdfsdfsdf`

	fmt.Println(synk)

	var litlerune = 'A'
	fmt.Println(litlerune)
	fmt.Printf("%c\n", litlerune)

	var zeroboo bool
	isprem := true
	nasSub := false
	fmt.Println(zeroboo, isprem, nasSub)
}
