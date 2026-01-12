package main

import (
	"errors"
	"fmt"
)

func convertCentToDollar(centVal int) {
	dollarVal := float64(centVal) / 100.00
	fmt.Printf("Transaction amount: $%.2f\n", dollarVal)
}

func getCurrRate(currCodes map[string]float64, code string) (float64, error) {
	rate, ok := currCodes[code]

	if !ok {
		return 0.0, errors.New("invalid currency code")
	}
	return rate, nil
}

func addCurr(currCodes map[string]float64, code string, rate float64) {
	currCodes[code] = rate
}

func processTransaction(amount float64) (float64, float64, error) {

	if amount <= 0 {
		return 0.0, 0.0, errors.New("invalid amount")
	}
	fee := amount * 0.029
	netAmount := amount - fee

	return netAmount, fee, nil

}

type Transaction struct {
	Amount float64
	Status bool
}

func approveTransaction(t *Transaction) {
	t.Status = true
}

func addFeeByValue(amount float64, fee float64) float64 {
	return amount + fee
}

func addFeeByPointer(amount *float64, fee float64) {
	*amount += fee
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

	// PROBLEM 3
	dailyTxns := []float64{45.99, 123.50, 8.75}
	dailyTxns = append(dailyTxns, 200.10, 15.30)
	sumTxns := 0.00
	numTxns := 0

	for _, txn := range dailyTxns {
		sumTxns += txn
		numTxns++
	}

	fmt.Println(sumTxns, numTxns)

	// PROBLEM 4
	codeMap := map[string]float64{
		"EUR": 1.08,
		"GBP": 1.27,
		"JPY": 0.0067,
	}

	rate_EUR, err := getCurrRate(codeMap, "EUR")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rate_EUR)
	}

	// PROBLEM 5
	type Merchant struct {
		ID       string
		Name     string
		Category string
		TxnCount int
	}

	m1 := Merchant{
		ID:       "amz",
		Name:     "Amazon",
		Category: "e-commerce",
		TxnCount: 100000,
	}

	m1.TxnCount += 500

	fmt.Println(m1)

	// PROBLEM 9
	txn1 := Transaction{
		Amount: 500.00,
		Status: false,
	}

	fmt.Println(txn1)
	approveTransaction(&txn1)
	fmt.Println(txn1)

	// PROBLEM 10
	amt = 100.00
	fee := 5.00

	totalAmt := addFeeByValue(amt, fee)
	fmt.Println(totalAmt)

	addFeeByPointer(&amt, fee)
	fmt.Println(amt)

}
