package main

import (
	"fmt"
)

func main() {
	var firstNum float64
	var lastNum float64
	var operation string

	for {
		fmt.Print("Enter First Number: ")
		fmt.Scanln(&firstNum)
		fmt.Println()
		fmt.Print("Enter Last Number: ")
		fmt.Scanln(&lastNum)
		fmt.Println()
		fmt.Println("Select operation:")
		fmt.Println("+", ": addition")
		fmt.Println("-", ": subtraction")
		fmt.Println("*", ": multiplication")
		fmt.Println("/", ": division")
		fmt.Println("  : quit")
		fmt.Println("  : help")
		fmt.Println()
		fmt.Scan(&operation)
		fmt.Println()

		if operation == "quit" {
			fmt.Println("Goodbye!!!")
			break
		}
		if operation == "help" {
			fmt.Println("This are the correct inputs:")
			fmt.Println("+ to add")
			fmt.Println("- to subtract")
			fmt.Println("* to multiply")
			fmt.Println("/ to divide")
		}

		switch operation {
		case "+":
			fmt.Println("add", firstNum, lastNum)
			fmt.Println("Result:", firstNum+lastNum)
			continue

		case "-":
			fmt.Println("sub", firstNum, lastNum)
			fmt.Println("Result:", firstNum-lastNum)
			continue

		case "*":
			fmt.Println("mul", firstNum, lastNum)
			fmt.Println("Result:", firstNum*lastNum)
			continue

		case "/":
			if lastNum == 0 {
				fmt.Println("Mathematical Error:")
				continue
			} else {
				fmt.Println("div", firstNum, lastNum)
				fmt.Println("Result:", firstNum/lastNum)
			}
		default:
			fmt.Println("Unknown command: Type 'help' for commands")
			continue
		}
	}

}

//		case "/":
/*if lastNum == 0 {
	fmt.Println("Mathematical Error:")
	continue
} else if firstNum < lastNum {
	fmt.Println("Mathematical Error")
	continue
} else if firstNum%lastNum == 0 {
	fmt.Println("div", firstNum, lastNum)
	fmt.Println("Result:", firstNum/lastNum)
	continue
} else if firstNum%lastNum != 0 {
	fmt.Println("Error: Not divisible")
	continue
}*/
