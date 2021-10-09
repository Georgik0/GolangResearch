package main

import "fmt"

func callMyName(name string) (string, error) {
	return name, nil
}

func main() {
	fmt.Println(callMyName("My name"))

	/* Отложенное выполнение */
	fmt.Println("\n--------------------Отложенное выполнение------------------------")
	{
		info := func(str string) string {
			fmt.Println(str)
			return "Out info"
		}

		defer func() {
			fmt.Println(info("Hi"))
		}()

		fmt.Println("After defer")
	}
}
