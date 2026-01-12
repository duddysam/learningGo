package main

import "fmt"

func convertCentToDollar(centVal int) {
	dollarVal := float64(centVal) / 100.00
	fmt.Printf("Transaction amount: $%.2f\n", dollarVal)
}

func main() {

	// PROBLEM 1
	var txnID string
	var amt float64
	var currCode string
	var mName string
	var isApproved bool

	txnID = "M101"
	amt = 100.01
	currCode = "USD"
	mName = "payPal"
	isApproved = true

	fmt.Printf("txn: %s of $%.2f, %s through %s, has status: %t\n", txnID, amt, currCode, mName, isApproved)

	// PROBLEM 2
	convertCentToDollar(1050)
}
