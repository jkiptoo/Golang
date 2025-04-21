package main

import (
	"errors"
	"fmt"
	"os"
)

/* ---- TASKS -----
- Validate user input.
- Show error message and exit if provided input is invalid.
- No negative numbers in input. \/
- No input of 0 \/
- Store calculated results in a file.
*/

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before Tax is: %.1f\n ", ebt)
	fmt.Printf("Profit is: %.1f\n ", profit)
	fmt.Printf("The Earnings before Tax to Profit Ratio is: %.1f\n", ratio)
	storeFinancials(ebt, profit, ratio)
}

func storeFinancials(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Earnings Before Tax: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("financialCalculations.txt", []byte(results), 0644)
}
func calculateFinancials(revenue, expenses, taxRate  float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Please enter positive values greater than 0.")

	}
	return userInput, nil
}
