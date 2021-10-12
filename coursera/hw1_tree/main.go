package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

func drawBranch(numTabs int, numPipe int, checkLast bool, name string) {
	for numTabs > 0 {
		fmt.Printf("\t")
		numTabs--
	}
	for numPipe > 0 {
		fmt.Printf("|\t")
		numPipe--
	}
	if checkLast {
		fmt.Printf("└───%v\n", name)
	} else {
		fmt.Printf("├───%v\n", name)
	}
}

func dirIsEmpty(path string, flag bool) bool {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if len(dir) != 0 && flag == true {
		return false
	}
	for _, file := range dir {
		if file.IsDir() {
			return false
		}
	}
	return true
}

func setCheckLast(iter int, len_dir int) bool {
	return iter == len_dir - 1
}

func drawTree(output *os.File, path string, flag bool, checkLast bool, prefix string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(dir, func (i, j int) bool {
		return dir[i].Name() < dir[j].Name()
	})

	var outTxt string = ""
	for iter, file := range dir {
		if flag == false && file.IsDir() == false {
			continue
		}
		//fmt.Println(file.Name(), "index: ", iter)
		checkLast = iter == len(dir) - 1
		if file.IsDir() {
			if dirIsEmpty(path + "/" + file.Name(), flag) {
				if checkLast {
					outTxt += prefix + "└───" + file.Name()
				} else {
					outTxt += prefix + "├───" + file.Name()
				}
				if flag == true {
					outTxt += " (" + strconv.Itoa(int(file.Size())) + "b)"
				}
				fmt.Fprintln(output, outTxt)
			} else {
				if checkLast {
					prefix += "\t"
				} else {
					prefix += "|\t"
				}
				drawTree(output, path + "/" + file.Name(), flag, checkLast, prefix)
			}
		} else {
			if checkLast {
				outTxt += prefix + "└───" + file.Name()
			} else {
				outTxt += prefix + "├───" + file.Name()
			}
			if flag == true {
				outTxt += " (" + strconv.Itoa(int(file.Size())) + "b)"
			}
			fmt.Fprintln(output, outTxt)
		}
		//fmt.Println(file.Name(), file.IsDir(), file.Size())
	}
	return nil
}

func dirTree(out *os.File, path string, flag bool) error {
	//lvl := 1
	drawTree(out, path, flag, true, "")
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
