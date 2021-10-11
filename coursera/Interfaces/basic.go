package Interfaces

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("No money in the wallet")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
}

//func main() {
//	myWallet := &Wallet{Cash: 100}
//	fmt.Printf("myWallet = %v\ttype: %[1]T\n", myWallet)
//	Buy(myWallet)
//	fmt.Println(myWallet)
//}
