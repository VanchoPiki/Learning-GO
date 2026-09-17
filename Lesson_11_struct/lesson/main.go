package main

import "fmt"

// 2
type Wallet struct {
	address  string
	balance  float64
	currency string
	isctive  bool
}

//4

type Wallet_1 struct {
	address string
	balance float64
}

func addmoney(w *Wallet_1, amount float64) {
	w.balance += amount
}

//5
type user struct {
	TGid int64
	username string
	Wallet Wallet_1
}

func main() {

	//1

	//type name_struct struct {
	//	pole_1 int
	//	pole_2 int
	//}

	//2
	w1 := Wallet{
		address:  "0x71c56g47g9",
		balance:  2.45,
		currency: "ETH",
		isctive:  true,
	}

	var w2 Wallet
	w2.address = "0x67f67f67f67"
	w2.currency = "TON"

	fmt.Println(w1, "\n", w2)

	//3

	w := Wallet{address: "0x12345", balance: 10.0}
	fmt.Println("Баланс кошелька: ", w.balance)

	w.balance = 15.5
	fmt.Println("Новый балик: ", w.balance)

	//4
	mywallet := Wallet_1{address: "0x12345", balance: 100.0}
	addmoney(&mywallet, 50.0)
	fmt.Println("Баланс пополнения: ", mywallet.balance)


	//5

	user := user{
		TGid : 12345678910,
		username: "ivan",
		Wallet: Wallet_1{
			address: "0x12345678910",
			balance: 500.0,
		},
	}

	fmt.Printf("%+v\n", user)
}
