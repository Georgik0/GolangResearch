package main

import (
	"fmt"

	"GolangResearch/research/research_imports/dir1"
	"GolangResearch/research/research_imports/dir2"
)

func init() {
	dir1.Config = 25
}

func main() {
	dir2.ChangeConfig()
	fmt.Println(dir1.Config)

	t1 := dir1.Tmp{A: 1}
	fmt.Println(t1.A)

	t2 := dir1.TmpLoc{}
	fmt.Println(t2)
}
