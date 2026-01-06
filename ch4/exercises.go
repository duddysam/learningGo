package main

import (
	"errors"
	"fmt"
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

func main() {

	formattedNum, err := formatCardNum("1234567890123456")

	fmt.Println(formattedNum, err)
}
