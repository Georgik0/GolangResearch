package main

import (
	"GolangResearch/coursera/Interfaces"
	"fmt"
)

func main() {
	myWallet := &Interfaces.Wallet{Cash: 100}
	fmt.Printf("myWallet = %v\ttype: %[1]T\n", myWallet)
	Interfaces.Buy(myWallet)
	fmt.Println(myWallet)

	BigWallet := &Interfaces.Wallet{Cash: 228}
	AppleWallet := Interfaces.ApplePay{Cash: 342}
	Interfaces.BuyCast(&AppleWallet)
	fmt.Printf("BigWallet:    %v\t%[1]T\t\n&BigWallet:    %v\t%[2]T\t\n", BigWallet, &BigWallet)
	fmt.Printf("AppleWallet:    %v\t%[1]T\t\n&AppleWallet:    %v\t%[2]T\t\n", AppleWallet, &AppleWallet)

	var someWallet Interfaces.Wallet = Interfaces.Wallet{100} // другой (полный) вариант объявления переменной
	fmt.Printf("\nRaw payment : %#v\n", someWallet)
	fmt.Printf("Raw payment : %s\n", someWallet)

	Interfaces.Buy_empty("string")
	Interfaces.Buy_empty(123)
	Interfaces.Buy_empty(someWallet)
	Interfaces.Buy_empty(AppleWallet)

	Interfaces.TestInterf()
}
