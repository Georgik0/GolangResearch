package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	reader := bufio.NewScanner(input)
	var prev string
	for reader.Scan() {
		txt := reader.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
