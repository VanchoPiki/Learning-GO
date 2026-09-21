package main

import "fmt"

//1

// 1. Интерфейс (Контракт)
type Payer interface {
	Pay(amount float64) bool
}

// 2. Первая структура: Биткоин-кошелек
type BitcoinWallet struct {
	BTCBalance float64
}

// Реализуем метод Pay для Bitcoin (просто пишем метод, никакого слова implements!)
func (b *BitcoinWallet) Pay(amount float64) bool {
	if b.BTCBalance < amount {
		fmt.Println("BTC: Недостаточно биткоинов!")
		return false
	}
	b.BTCBalance -= amount
	fmt.Printf("BTC: Успешно оплачено %.4f BTC\n", amount)
	return true
}

// 3. Вторая структура: Обычная банковская карта
type CreditCard struct {
	CardNumber string
	Limit      float64
}

// Реализуем метод Pay для карты
func (c *CreditCard) Pay(amount float64) bool {
	if c.Limit < amount {
		fmt.Println("Банк: Превышен лимит по карте!")
		return false
	}
	c.Limit -= amount
	fmt.Printf("Банк: Оплачено картой %s на сумму %.2f руб.\n", c.CardNumber, amount)
	return true

}

//2

// Функция принимает ЛЮБОГО, кто удовлетворяет интерфейсу Payer
func BuyCourse(p Payer, price float64) {
	fmt.Println("Попытка провести оплату...")
	p.Pay(price) // Вызывается метод Pay той структуры, которую передали!
}

func main() {
	myBTC := &BitcoinWallet{BTCBalance: 1.5}
	myCard := &CreditCard{CardNumber: "4400...1234", Limit: 5000.0}

	// Передаем разные структуры в одну и ту же функцию:
	BuyCourse(myBTC, 0.2)   // Сработает логика биткоина!
	BuyCourse(myCard, 2000) // Сработает логика банковской карты!

	//3

	myBTC_2 := &BitcoinWallet{BTCBalance: 1.5}
	myCard_2 := &CreditCard{CardNumber: "4400...1234", Limit: 5000.0}

	// Срез содержит разные типы данных, объединенные одним интерфейсом:
	paymentMethods := []Payer{myBTC_2, myCard_2}

	// Списываем по 100 условных единиц со всех способов оплаты:
	for _, method := range paymentMethods {
		method.Pay(100)
	}
}
