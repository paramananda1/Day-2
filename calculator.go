package main

import "fmt"


func Add(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) float64 {
	return num1 / num2
}

func calculator() {
	var operation string
	fmt.Println("Please select an operation: +, -, *, / : ")
	fmt.Scanln(&operation)
	var num1 int
	fmt.Println("Please input the first number: ")
	fmt.Scanf("%d", &num1)
	var num2 int
	fmt.Println("Please input the Second number: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Result:  ")
	switch operation {
	case "+":
		fmt.Println( Add(num1, num2))
	case "-":
		fmt.Println( Subtract(num1, num2))

	case "*":
		fmt.Println(Multiply(float64(num1), float64(num2)))

	case "/":
		fmt.Println(Divide(float64(num1), float64(num2)))

	default:
		fmt.Println("Invalid operation selected. Please try again!")
	}

}

