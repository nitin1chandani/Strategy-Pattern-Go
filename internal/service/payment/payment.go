package payment

import "fmt"

type PaymentGateway interface {
	Pay(amount string)
}

type UPI struct{}

func (u UPI) Pay(amount string) {
	fmt.Println("Payment done")
}

type Card struct{}

func (c Card) Pay(amount string) {
	fmt.Println("Payment done")
}

type NetBanking struct{}

func (nb NetBanking) Pay(amount string) {
	fmt.Println("Payment done")
}

type Pay struct {
	paymentGateway PaymentGateway
}

func (p Pay) PayNow(amount string) {
	p.paymentGateway.Pay(amount)
}
