package main

import "fmt"

// 1
type SimpleWallet struct {
	Address string
	Balance float64
}

func (s *SimpleWallet) PrintBalance() {
	fmt.Printf("Кошелек [%s] | Баланс: [%.2f]\n", s.Address, s.Balance)
}

type PremiumWallet struct {
	SimpleWallet
	CashBack float64
}

func (p *PremiumWallet) PrintBalance() {
	fmt.Printf("[PREMIUM VIP] Кошелек [%s] | Баланс: [%.2f] | Кешбэк: [%.0f]%%\n", p.Address, p.Balance, p.CashBack*100)
}

func main() {
	prmwal := PremiumWallet{
		SimpleWallet{
			"0xw123123",
			123.4,
		},
		0.005,
	}

	fmt.Println(prmwal.Address)
	prmwal.PrintBalance()
	prmwal.SimpleWallet.PrintBalance()
}
