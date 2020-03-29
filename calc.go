package main

import (
	"fmt"
)

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) float64 {
	return num1 / num2
}

func main() {
	var operation string
	var num1, num2 float64

enterValues:

	fmt.Println("Please enter fist num:")
	fmt.Scanln(&num1)

	fmt.Println("Please enter last num:")
	fmt.Scanln(&num2)

	fmt.Println("Please select an operation: +, -, *, /:")
	fmt.Scanln(&operation)

	fmt.Println("Result:")

	switch operation {
	case "+":
		fmt.Println(Add(num1, num2))
	case "-":
		fmt.Println(Subtract(num1, num2))
	case "/":
		fmt.Println(Divide(num1, num2))
	case "*":
		fmt.Println(Multiply(num1, num2))
	default:
		fmt.Println("Invalid operation selected. plz try again!")
		goto enterValues
	}

}
