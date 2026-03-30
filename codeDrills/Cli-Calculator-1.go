package main

import "fmt"

func firstProject() {
	var firstNum int
	var lastNum int
	var operation string

	for {
		fmt.Print("Enter First Number: ")
		fmt.Scanln(&firstNum)
		if firstNum <= 0 || firstNum >= 9 {
			fmt.Println("Error: Wrong input.")
			continue
		}
		fmt.Println()
		fmt.Print("Enter Last Number: ")
		fmt.Scanln(&lastNum)
		/*if lastNum <= 0 || lastNum >= 9 {
		    fmt.Println("Error: Wrong Input.")
		    continue
		}*/
		fmt.Println()
		fmt.Println("Choose operation: ")
		fmt.Println()
		fmt.Println("add", firstNum, lastNum)
		fmt.Println("sub", firstNum, lastNum)
		fmt.Println("mul", firstNum, lastNum)
		fmt.Println("div", firstNum, lastNum)
		fmt.Println()
		fmt.Scanln(&operation)
		fmt.Println()

		switch operation {
		case "add":
			fmt.Println("Result:", firstNum+lastNum)
			return

		case "sub":
			fmt.Println("Result:", firstNum-lastNum)
			return

		case "mul":
			fmt.Println("Result:", firstNum*lastNum)
			return

		case "div":

			fmt.Println("Result:", firstNum/lastNum)
			return

		}
	}
}
