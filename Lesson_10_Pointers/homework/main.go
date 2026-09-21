package main

import "fmt"

// 2
func double(n *int) {
	*n = *n * 2
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("\n//1\n")

	count := 10
	ptr := &count
	fmt.Println(ptr)
	*ptr = 50
	fmt.Println(count)

	fmt.Println("\n//2\n")

	n := 2
	double(&n)
	fmt.Println(n)

	fmt.Println("\n//3\n")

	a := 2
	b := 3
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Println(&n)
	fmt.Println()
}
