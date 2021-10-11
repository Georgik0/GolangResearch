package Interfaces

import "fmt"

type ApplePay struct {
	Cash int
}

func (ap *ApplePay) Pay(amount int) error {
	if ap.Cash < amount {
		return fmt.Errorf("No money in the wallet")
	}
	ap.Cash -= amount
	return nil
}

func BuyCast(p Payer) {
	switch p.(type) {
	case *ApplePay:
		fmt.Printf("Payer ApplePay: %T\t%[1]v", p)
	case *Wallet:
		fmt.Printf("Payer Wallet: %T\t%[1]v", p)
	default:
		fmt.Println("Default")
	}
}
