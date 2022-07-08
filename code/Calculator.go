package main

import (
	"fmt"
)

func main() {
	printPepit("Indica la operación")
	var op string
	fmt.Scan(&op)

	if op != "+" || op != "-" || op != "*" || op != "x" || op != "/" {
		printPepit("Error! selecciona +, -, * o x, /")
		fmt.Scan(&op)
	}

	printPepit("Inserta dos numeros")

	var num1 float64
	fmt.Scan(&num1)

	var num2 float64
	fmt.Scan(&num2)

	if op == "+" {
		fmt.Println(num1 + num2)
	} else if op == "-" {
		fmt.Println(num1 - num2)
	} else if op == "*" || op == "x" {
		fmt.Println(num1 * num2)
	} else if op == "/" {
		fmt.Println(num1 / num2)
	}
}

func printPepit(str string) {
	fmt.Println(str)
}
