package Interfaces

import (
	"fmt"
)

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

func Buy_empty(in interface{}) {
	var p Payer
	var ok bool
	fmt.Printf("p: %T, %[1]v\n", p)
	//p = &Wallet{100}
	if p, ok = in.(Payer); !ok {
		fmt.Printf("ok: %v; %T не является интерфейсом %v\n\n", ok, in, p)
	} else {
		fmt.Printf("%T интерфейс Payer\n\n", in)
	}
}

func TestInterf() {
	var emptyItf Payer
	var wallet = Wallet{23}
	emptyItf = &Wallet{100}
	fmt.Printf("empty_itf: %v   %[1]T\n", emptyItf)
	emptyItf = &wallet
	fmt.Printf("empty_itf: %v   %[1]T", emptyItf)
}
