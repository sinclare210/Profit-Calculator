package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var taxRate float64
	// var expenses float64
	// var revenue float64

	// fmt.Print("What are your expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := userInput("What are your expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("What is your revenue: ")
	// fmt.Scan(&revenue)

	revenue, err := userInput("What are your revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := userInput("What is your taxt rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningBeforeTax, earningAfterTax, profitRatio := result(expenses, revenue, taxRate)
	writeResultToFile(earningBeforeTax, earningAfterTax, profitRatio)

	fmt.Printf("Your Earning Before Tax is %0.2f\n", earningBeforeTax)

	fmt.Printf("Your Earning After Tax is %0.2f\n", earningAfterTax)

	fmt.Printf("Your Profit Ratio is %0.2f\n", profitRatio)
}

func userInput(input string) (float64, error) {
	var userInput float64
	fmt.Print(input)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Can't input a value less than or equal to zeroN")
	}
	return userInput, nil

}

func result(expenses, revenue, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses

	taxAmount := earningBeforeTax * taxRate

	earningAfterTax := earningBeforeTax - taxAmount

	profitRatio := (earningAfterTax / revenue) * 100

	return earningBeforeTax, earningAfterTax, profitRatio
}

func writeResultToFile(ebt, profit, ratio float64) {
	resultText := fmt.Sprintf("EBT: %.1f\nProfit:%.1f\nRatio:%.1f", ebt, profit, ratio)
	os.WriteFile("result.text", []byte(resultText), 0644)
}
