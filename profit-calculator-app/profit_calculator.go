package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64

	revenue, err := getUserInput("Input Revenue: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic(err)
		return
	}

	expenses, err = getUserInput("Input Expenses: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic(err)
		return
	}

	taxRate, err = getUserInput("Input Tax Rate: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings Before Tax: ", ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, expenses, ratio
}

func getUserInput(input_text string) (float64, error) {
	var input_value float64
	fmt.Print(input_text)
	fmt.Scan(&input_value)

	if input_value <= 0 {
		return 0, errors.New("Please input value more than 0")
	}

	return input_value, nil
}
