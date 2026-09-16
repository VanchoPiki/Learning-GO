package main

import "fmt"

func main() {
	//1
	fmt.Println("\n//1\n")

	phonebook := make(map[string]string)

	phonebook["Мама"] = "+79999997777"
	phonebook["Друг"] = "+71231232323"
	phonebook["Работа"] = "+71231231212"
	fmt.Println(phonebook)
	phonebook["Друг"] = "+7 999 999 12 12"
	fmt.Println(phonebook)
	delete(phonebook, "Работа")
	fmt.Println(phonebook)

	//2
	fmt.Println("\n//2\n")

	permissions := map[string]bool{
		"alex":  true,
		"john":  false,
		"guest": false,
	}

	fmt.Println(permissions)

	if name, ok := permissions["alex"]; ok {
		if name == true {
			fmt.Println("Доступ разрешен")
		} else {
			fmt.Println("Доступ запрещен")
		}
	} else {
		fmt.Println("Пользователь не найден")
	}

	//3
	fmt.Println("\n//3\n")

	cart := map[string]float64{
		"Хлеб":    40.0,
		"Молоко":  85.5,
		"Сыр":     240.0,
		"Шоколад": 95.0,
	}
	total := 0.0
	index := 0
	for key, value := range cart {
		total += value
		index++
		fmt.Printf("%d. %s - %.2f\n", index, key, value)
	}
	fmt.Println("---------------\nОбщий чек -", total)
}
