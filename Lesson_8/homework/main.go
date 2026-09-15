package main

import "fmt"

func main() {

	fmt.Println("\n//1\n")

	var todos []string

	todos = append(todos, "Проснуться")
	todos = append(todos, "Позавтракать", "Поучить Go", "Пойти гулять")

	for i, work := range todos {
		fmt.Printf("%d. %s\n", i+1, work)
	}

	fmt.Println("\n//2\n")

	numbers := make([]int, 0, 3)

	for i := 0; i < 5; i++ {
		numbers = append(numbers, i+12)
		fmt.Printf("Элемент: %d | len: %d | cap: %d\n", numbers[i], len(numbers), cap(numbers))
	}

	fmt.Println("\n//3\n")

	nums := [6]int{10, 20, 30, 40, 50, 60}

	slicePart := nums[2:5]

	slicePart[0] = 999

	fmt.Println(nums, "\n", slicePart, "\n")

	fmt.Println("\n//4\n")

	team1 := []string{"Алексей", "Иван"}
	team2 := []string{"Ольга", "Мария"}

	allTeams := append(team1, team2...)

	fmt.Println(allTeams, len(allTeams))

}
