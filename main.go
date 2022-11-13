package main

import "fmt"

func main() {
	var firstNumber int
	print("first number: ")
	fmt.Scan(&firstNumber)

	var operation string
	print("operation: ")
	fmt.Scan(&operation)

	var secondNumber int
	print("second number: ")
	fmt.Scan(&secondNumber)

	if(operation == "+") {
		print("result : ")
		println(firstNumber + secondNumber)
	}
	if(operation == "-") {
		print("result : ")
		println(firstNumber - secondNumber)
	}
	if(operation == "*") {
		print("result : ")
		println(firstNumber * secondNumber)
	}
	if(operation == "/") {
		print("result : ")
		println(firstNumber / secondNumber)
	}

}