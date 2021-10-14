package main

import (
	"fmt"
	"io"
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
		fmt.Printf("│\t")
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

func getSize(size int64, dirEmptyFlag bool) string {
	if size == 0 || dirEmptyFlag {
		return "empty"
	}
	return strconv.Itoa(int(size)) + "b"
}

func getOutTxt(file os.FileInfo, checkLast bool, flag bool, prefix string, dirEmptyFlag bool) string {
	var outTxt = ""
	if checkLast {
		outTxt += prefix + "└───" + file.Name()
	} else {
		outTxt += prefix + "├───" + file.Name()
	}
	if flag == true {
		size := getSize(file.Size(), dirEmptyFlag)
		outTxt += " (" + size + ")"
	}
	return outTxt
}

func processingDir(output io.Writer, path string, checkLast bool, file os.FileInfo, flag bool, prefix string) {
	var outTxt = ""
	if dirEmptyFlag := dirIsEmpty(path + "/" + file.Name(), flag); dirEmptyFlag {
		outTxt := getOutTxt(file, checkLast, flag, prefix, dirEmptyFlag)
		fmt.Fprintln(output, outTxt)
	} else {
		if checkLast {
			outTxt += prefix + "└───" + file.Name()
			prefix += "\t"
		} else {
			outTxt += prefix + "├───" + file.Name()
			prefix += "│\t"
		}
		fmt.Fprintln(output, outTxt)
		drawTree(output, path + "/" + file.Name(), flag, checkLast, prefix)
	}
}

func getCheckLast(flag bool, iter int, dir []os.FileInfo) bool {
	iter++
	if flag == false {
		for ; iter < len(dir); iter++ {
			if dir[iter].IsDir() {
				return false
			}
		}
		return true
	}
	return iter == len(dir)
}

func drawTree(output io.Writer, path string, flag bool, checkLast bool, prefix string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(dir, func (i, j int) bool {
		return dir[i].Name() < dir[j].Name()
	})

	for iter, file := range dir {
		if flag == false && file.IsDir() == false {
			continue
		}
		checkLast = getCheckLast(flag, iter, dir)
		if file.IsDir() {
			processingDir(output, path, checkLast, file, flag, prefix)
		} else {
			outTxt := getOutTxt(file, checkLast, flag, prefix, false)
			fmt.Fprintln(output, outTxt)
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, flag bool) error {
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
