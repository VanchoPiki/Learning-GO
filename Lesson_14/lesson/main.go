package main

import "fmt"

// 1
type Wallet struct {
	Balance float64
}

// 1. Обычная вложенность (мы это уже делали):
type RegularUser struct {
	Wallet Wallet // Поле имеет имя 'Wallet'
}

// Чтобы добраться до баланса: user.Wallet.Balance

// 2. ВСТРАИВАНИЕ (Embedding):
type SuperUser struct {
	Wallet // ИМЕНИ НЕТ! Просто указали тип структуры
}

// Магия: user.Balance (поле само всплыло наверх!)

//2

// Базовая "деталь LEGO": Обычный пользователь
type User struct {
	Username string
	Email    string
}

// Метод обычного пользователя
func (u *User) Login() {
	fmt.Printf("Пользователь %s вошел в систему\n", u.Username)
}

// Собираем Админа: Админ — это User + права модератора
type Admin struct {
	User       // ВСТРАИВАНИЕ! (Админ получает все поля и методы User)
	SuperPower string
}

// 3
// Переопределяем Login для Admin
func (a *Admin) Login() {
	fmt.Printf("⚠️ АДМИНИСТРАТОР %s вошел с повышенными правами!\n", a.Username)
}

//4

type SecuritySystem struct {
	Has2FA bool
}

type BankAccount struct {
	Balance float64
}

// VIP-клиент собран из двух разных сущностей:
type VIPUser struct {
	SecuritySystem // встроен раз
	BankAccount    // встроен два
	Tier           int
}

func main() {
	//2
	// Создаем админа:
	admin := Admin{
		User: User{
			Username: "boss_alex",
			Email:    "alex@crypto.com",
		},
		SuperPower: "Бан навсегда",
	}

	// 1. Поля всплыли наверх (не нужно писать admin.User.Username):
	fmt.Println("Имя админа:", admin.Username)

	// 2. Методы тоже всплыли наверх!
	// У Admin нет своего метода Login(), но он берет его у встроенного User:
	admin.Login() // Выведет: Пользователь boss_alex вошел в систему

	//3

	admin1 := Admin{
		User: User{Username: "boss_alex"},
	}

	// Вызовется переопределенный метод админа:
	admin1.Login()

	// Но если очень надо, можно вызвать и старый базовый метод:
	admin1.User.Login()
}
