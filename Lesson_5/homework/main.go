package main

import "fmt"

// 1
func rublesToDollars(rubles, rate float64) float64 {
	dollars := rubles / rate
	return dollars
}

// 2
func max_num(a, b int) (int, string) {
	switch {
	case a > b:
		return a, "Больше"
	case a < b:
		return b, "Больше"
	default:
		return a * b * 0, "Они одинаковы"
	}
}

// 3
func rectangleProps(width, height float64) (float64, float64) {
	perim := 2 * (width + height)
	plosh := width * height

	return perim, plosh
}

// 4
func safeDivide(a, b float64) (float64, bool) {
	switch {
	case b == 0:
		return 0, false
	default:
		return a / b, true
	}
}

func main() {
	//1
	dollars := rublesToDollars(99.0, 92.5)
	fmt.Printf("Ваши рубли после конвертации: %.2f\n", dollars)

	//2
	big_num, text := max_num(3, 6)
	fmt.Printf("%d %s\n", big_num, text)

	//3
	perim, plosh := rectangleProps(3.0, 4.0)
	fmt.Printf("Периметр: %.0f, Площадь: %.2f\n", perim, plosh)

	//4
	delit, mojhn := safeDivide(4.5, 5.5)
	fmt.Printf("%.2f, %t\n", delit, mojhn)

	if delit1, hz := safeDivide(4.5, 0) ; !hz {
		fmt.Println("Нельзя делить на 0")
	} else {
		fmt.Printf("Результат: %.0f\n", delit1)
	}
}
