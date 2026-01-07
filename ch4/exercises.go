package main

import (
	"errors"
	"fmt"
	"unicode"
)

func formatCardNum(cardNum string) (string, error) {
	if len(cardNum) != 16 {
		return "", errors.New("Invalid card number")
	}
	out := ""
	for i, r := range cardNum {
		if i > 0 && i%4 == 0 {
			out += " "
		}
		out += string(r)
	}
	return out, nil

}

func nameValidator(name string) (bool, error) {

	nameRunes := []rune(name)

	if len(nameRunes) < 2 || len(nameRunes) > 50 {
		return false, errors.New("Invalid name length")
	}

	for i, r := range nameRunes {

		if !unicode.IsLetter(r) && (r != ' ' && r != '\'' && r != '-') {
			return false, errors.New("Invalid Character:")
		}
		if (i == 0 || i == len(nameRunes)-1) && r == ' ' {
			return false, errors.New("Cannot start or end with a space")
		}
	}

	return true, nil
}

func filterTxns(txnAmounts []float64, threshold float64) ([]float64, int) {

	out := make([]float64, 0, len(txnAmounts))
	removed := 0

	for _, v := range txnAmounts {

		if v > threshold {
			out = append(out, v)
		} else {
			removed++
		}
	}

	return out, removed

}

func main() {

	// TEST PROBLEM 1
	formattedNum, err := formatCardNum("1234567890123456")

	fmt.Println(formattedNum, err)

	// TEST PROBLEM 2
	valid, err := nameValidator("Sam Duddy")
	fmt.Println("true ==", valid, err)

	valid, err = nameValidator(" Sam Duddy")
	fmt.Println("false ==", valid, err)

	valid, err = nameValidator("Sam Duddy ")
	fmt.Println("false ==", valid, err)

	valid, err = nameValidator("Sam Duddy's")
	fmt.Println("true ==", valid, err)

	valid, err = nameValidator("Sam Duddy-s")
	fmt.Println("true ==", valid, err)

	valid, err = nameValidator("Sam Duddy$s")
	fmt.Println("false ==", valid, err)

	// PROBLEM 3
	txns := []float64{100.00, 101.00, 400.00, 99.00, 73.00, 62.00}

	filteredTxns, removed := filterTxns(txns, 75.00)
	fmt.Println(filteredTxns, removed)

}
