package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calc(i int, j int, equation []string) string {
	equation = equation[i+1 : j-1]
	fmt.Println(equation)
	for i = 0; i <= len(equation); i = i + 2 {
		//convert string to num type for calculation
		num1, err := strconv.Atoi(equation[i])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(equation[i+2])
		if err != nil {
			log.Fatal(err)
		}
		switch equation[i+1] {
		case "*":
			equation[i+2] = strconv.Itoa(num1 * num2)
		case "/":
			equation[i+2] = strconv.Itoa(num1 / num2)
		case "+":
			equation[i+2] = strconv.Itoa(num1 + num2)
		case "-":
			equation[i+2] = strconv.Itoa(num1 - num2)
		}
	}
	return equation[i+2]
}

func shiftandcalc(i int, j int, equation []string) []string {

	// if string haven't any brackets we just calculate string and return
	if i == 0 && j == 0 {
		calc(i, j, equation)
		return equation
	}
	//calc string between brackets
	equation[i] = calc(i+1, j-1, equation)
	for j < len(equation) {
		//shift array
		equation[i+1] = equation[j+1]
		j++
	}
	return equation
}

func main() {
	var sequation int
	var i, j int = 0, 0

	//reade Question string
	fmt.Println("Please enter your equation: (space after any argument)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	//fmt.Println(line)

	//convert string to Array
	equation := strings.Fields(line)
	//fmt.Println(equation)

	//find "("
	for i = 0; i < len(equation); i++ {

		fmt.Println("i in first loop:", i) //check i num in first loop
		fmt.Println(equation[i])

		if equation[i] == `(` {
			fmt.Println("ok")

			for j = i + 1; j < len(equation); j++ {
				fmt.Println("i in second loop:", i, j)

				//find ")"
				if equation[j] == `(` {
					i = sequation
					i = j
					fmt.Println("state (")
					continue

					//find ")"
				} else if equation[j] == `)` {
					equation = shiftandcalc(i, j, equation)
					i = sequation
					fmt.Println("state )")

				} // else {
				//fmt.Println("Please close Brackets")
				//break
				//}
			}
		}
	}
	fmt.Println("Your answer is:", calc(i, j, equation))
}
