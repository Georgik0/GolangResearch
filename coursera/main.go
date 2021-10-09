package main

import (
	"GolangResearch/coursera/Animal"
	"GolangResearch/coursera/Car"
	"fmt"
)

func main() {
	var bmw Car.My_Car
	bmw.Public_name = "BMW"
	fmt.Println(bmw)

	var myDog = Animal.Dog{}
	fmt.Println(myDog)
}
