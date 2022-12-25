package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	if err := setIDFromPart("elektronika-15500"); err != nil {
		fmt.Println(err)
	}
}

func setIDFromPart(part string) error {
	runes := []rune(part)

	if len(runes) == 0 {
		return errors.New("error converting path to runes")
	}

	// go from last rune to first rune until we find non digit rune or reach index 0
	// after this loop idStartIndex fill contain the first index for digit rune
	// also have to check that last rune must be digit
	var idStartIndex = len(runes) - 1
	if !unicode.IsDigit(runes[idStartIndex]) {
		return errors.New("error converting path to runes")
	}

	for idStartIndex > 0 && unicode.IsDigit(runes[idStartIndex-1]) {
		idStartIndex--
	}

	//now take a slice of runes from idStartIndex to the end
	digitRunes := runes[idStartIndex:]
	// convert digits to string again and parse integer
	id, err := strconv.ParseInt(string(digitRunes), 10, 64)
	if err != nil {
		return errors.New("error parseInt")
	}

	fmt.Println(id)
	return nil
}
