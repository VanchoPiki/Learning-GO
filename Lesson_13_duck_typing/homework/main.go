package main

import "fmt"

//1

type Notificer interface {
	Notify(message string)
}

//2

type TelegramNotificer struct {
	Username string
}

func (u *TelegramNotificer) Notify(message string) {
	fmt.Printf("[Teleegram -> %s]: %s\n", u.Username, message)
}

type EmailNotificer struct {
	Email string
}

func (e *EmailNotificer) Notify(message string) {
	fmt.Printf("[Email -> %s]: %s\n", e.Email, message)
}

//3

func SendAlert(n Notificer, text string) {
	n.Notify(text)
}

func main() {

	//3
	tgnoc := TelegramNotificer{Username: "@sex"}
	emnoc := EmailNotificer{Email: "sex@gmailo.com"}

	SendAlert(&tgnoc, "Вам начислен 1 BTC!")
	SendAlert(&emnoc, "Вам начислен 1 BTC!")

	//4

	notification := []Notificer{&tgnoc, &emnoc}

	for _, method := range notification {
		method.Notify("Сервер уходит на техобслуживание!")
	}
}
