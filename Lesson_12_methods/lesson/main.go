package main

import "fmt"

type Wallet struct {
	Address string
	Balance float64
}

// 1. Read-Only метод (Value receiver): просто выводит инфу
func (w Wallet) PrintInfo() {
	fmt.Printf("Кошелек %s | Баланс: %.2f\n", w.Address, w.Balance)
}

// 2. Метод с изменением данных (Pointer receiver): пополняет баланс
func (w *Wallet) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return
	}
	w.Balance += amount
	fmt.Printf("Успешно пополнено на %.2f\n", amount)
}

// 3. Метод со списанием: возвращает bool (хватило денег или нет)
func (w *Wallet) Withdraw(amount float64) bool {
	if amount > w.Balance {
		fmt.Println("Недостаточно средств!")
		return false
	}
	w.Balance -= amount
	return true
}

func main() {
	// Создаем кошелек
	myWallet := Wallet{Address: "0x777...AAA", Balance: 100.0}

	// Вызываем методы через точку!
	myWallet.PrintInfo() // Кошелек 0x777...AAA | Баланс: 100.00

	myWallet.Deposit(50.0)
	myWallet.PrintInfo() // Баланс стал 150.00!

	// Пробуем списать деньги
	success := myWallet.Withdraw(70.0)
	if success {
		fmt.Println("Списание прошло успешно!")
	}

	myWallet.PrintInfo() // Баланс стал 80.00!
}
