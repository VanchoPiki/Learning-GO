package main

import "fmt"

func main() {
	//1

	numbers := []int{1, 2, 3} // Это СРЕЗ (длина 3, емкость 3)
	fmt.Println(numbers)

	//2

	// Срез длиной 3 и емкостью 5, заполнен нулями [0, 0, 0]
	s := make([]int, 3, 5)
	fmt.Println(s)

	//3

	months := [5]string{"Янв", "Фев", "Мар", "Апр", "Май"}

	spring := months[2:4] // Возьмет индексы 2 и 3 ("Мар", "Апр")
	fmt.Println(spring)   // [Мар Апр]

	//4

	// Длина 3, Емкость 5
	nums := make([]int, 3, 5)

	fmt.Printf("Длина: %d, Емкость: %d, Значения: %v\n", len(nums), cap(nums), nums)
	// Выведет: Длина: 3, Емкость: 5, Значения: [0 0 0]

	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	// nums[3] = 40 // ОШИБКА! По индексу можно обращаться только в пределах len (0, 1, 2)!

	//5

	var list []string // Пустой nil-срез (len = 0, cap = 0)

	// Добавляем один элемент
	list = append(list, "Кофе")

	// Добавляем сразу несколько элементов через запятую
	list = append(list, "Чай", "Сахар")

	fmt.Println(list) // [Кофе Чай Сахар]

	//6

	t := make([]int, 2, 2)                           // len: 2, cap: 2 (мест нет!)
	fmt.Printf("len: %d, cap: %d\n", len(t), cap(t)) // len: 2, cap: 2

	t = append(t, 99) // ПРЕВЫШЕНИЕ ЕМКОСТИ!

	// Go увеличил емкость в 2 раза!
	fmt.Printf("len: %d, cap: %d\n", len(t), cap(t)) // len: 3, cap: 4

	//7

	a := []int{1, 2}
	b := []int{3, 4, 5}

	// Распаковываем срез 'b' по элементам:
	a = append(a, b...)
	fmt.Println(a) // [1 2 3 4 5]

	//8

	base := []int{1, 2, 3, 4}
	sub := base[1:3] // [2, 3]

	sub[0] = 999 // Меняем в sub

	fmt.Println("sub: ", sub)  // [999, 3]
	fmt.Println("base:", base) // [1, 999, 3, 4] — ОРИГИНАЛ ТОЖЕ ИЗМЕНИЛСЯ!
}
