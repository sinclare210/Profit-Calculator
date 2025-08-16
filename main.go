package main

import (
	"fmt"
)

func main (){
	var taxRate float64
	var expenses float64
	var revenue float64

	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	fmt.Print("What are your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("What is your revenue: ")
	fmt.Scan(&revenue)

	earningBeforeTax := revenue - expenses

	taxAmount := earningBeforeTax * taxRate

	earningAfterTax := earningBeforeTax - taxAmount

	profitRatio := (earningAfterTax/revenue) * 100

	fmt.Println("Your Earning Before Tax is ", earningBeforeTax)

	fmt.Println("Your Earning After Tax is ", earningAfterTax)

	fmt.Println("Your Profit Ratio is ", profitRatio)
}