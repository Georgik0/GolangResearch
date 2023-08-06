package main

import "GolangResearch/research/nil/nil-receiver/test1"

func main() {

}

func f1() {
	t1 := &test1.Rec{A: 1}
	t1.Check()

	t2 := interface{}(nil)
}
