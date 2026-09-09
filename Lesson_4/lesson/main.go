package main

import "fmt"

var totalfilm = 0

const zzz = 10000

func addfilm() {
	totalfilm++
}
func main() {
	addfilm()
	addfilm()

	fmt.Println(totalfilm)
}
