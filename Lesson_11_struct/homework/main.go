package main

import (
	"fmt"
)

// 1
type Transaction struct {
	from    string
	to      string
	amount  float64
	success bool
}

//2

type Wallet struct {
	Address string
	Balance float64
}

func deposit(w *Wallet, amount float64) {
	if amount <= 0 {
		fmt.Println("сумма пополнения должна быть больше 0")
	} else {
		w.Balance += amount
	}
}

//3

type TGuser struct {
	ID       int64
	username string
	wallet   Wallet
}

func main() {

	//1

	transac := Transaction{
		from:    "0x123456",
		to:      "0x78910",
		amount:  167.892949,
		success: true,
	}

	fmt.Printf("%+v\n", transac)

	//2
	w := Wallet{
		Address: "0x123456",
		Balance: 10.0,
	}

	deposit(&w, 25.5)

	fmt.Println(w.Balance)

	//3

	user_tg := TGuser{
		ID:       119233231235,
		username: "Kallevvann",
		wallet: Wallet{
			Address: "0x123456",
			Balance: 10.0,
		},
	}

	fmt.Printf("Юзер %s с ID %d имеет баланс: %.1f\n", user_tg.username, user_tg.ID, user_tg.wallet.Balance)

	//4

	wallets := []Wallet{
		{Address: "0x123456", Balance: 10.0},
		{Address: "0x123123", Balance: 15.0},
		{Address: "0x342342", Balance: 90.0},
	}

	total := 0.0

	for _, amount := range wallets {
		total += amount.Balance
	}

	fmt.Println(total)
}
