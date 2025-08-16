package main

import (
	"fmt"
)

func main (){
	// var taxRate float64
	// var expenses float64
	// var revenue float64


	// fmt.Print("What are your expenses: ")
	// fmt.Scan(&expenses)
	expenses := userInput("What are your expenses: ")

	// fmt.Print("What is your revenue: ")
	// fmt.Scan(&revenue)

	revenue := userInput("What are your revenue: ")

	taxRate := userInput("What is your taxt rate: ")

	earningBeforeTax,earningAfterTax,profitRatio := result(expenses,revenue,taxRate)

	fmt.Println("Your Earning Before Tax is ", earningBeforeTax)

	fmt.Println("Your Earning After Tax is ", earningAfterTax)

	fmt.Println("Your Profit Ratio is ", profitRatio)
}

func userInput (input string) float64{
	var userInput float64
	fmt.Print(input)
	fmt.Scan(&userInput)
	return  userInput
}

func result (expenses,revenue,taxRate float64)(float64, float64, float64){
	earningBeforeTax := revenue - expenses

	taxAmount := earningBeforeTax * taxRate

	earningAfterTax := earningBeforeTax - taxAmount

	profitRatio := (earningAfterTax/revenue) * 100

	return earningBeforeTax,earningAfterTax,profitRatio
}