package main

import "fmt"

type TGUser struct {
	ID       int64
	Username string
	Balance  float64
}

//1

func (w TGUser) PrintProfile() {
	fmt.Printf("Пользователь: %s | Баланс: %.2f USDT\n", w.Username, w.Balance)
}

// 2
func (u *TGUser) Withdraw(amount float64) bool {
	if amount > u.Balance {
		fmt.Println("Недостатч. средств.\n")
		return false
	}
	u.Balance -= amount
	fmt.Printf("Удачно! %.2f\n", u.Balance)
	return true
}

// 3
func (u *TGUser) Transfer(to *TGUser, amount float64) bool {
	if u.Balance < amount {
		fmt.Println("Недостатч. средств.\n")
		return false
	}
	u.Balance -= amount
	to.Balance += amount
	fmt.Printf("Удачно! %.2f\n", u.Balance)
	return true
}

func main() {
	//1
	tgacc := TGUser{1231212123, "sdfssfsdf", 123123.12312}
	tgacc.PrintProfile()
	//2
	amount := 123.123
	tgacc.Withdraw(amount)
	//3
	alice := TGUser{987654321, "@alice", 100.0}
	bob := TGUser{123456789, "@bob", 10.0}
	alice.PrintProfile()
	bob.PrintProfile()

	alice.Transfer(&bob, 40.0)
	alice.PrintProfile()
	bob.PrintProfile()
}
