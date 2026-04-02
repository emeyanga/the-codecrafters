package main

import (
	"fmt"
)

func main() {
	var firstNum float64
	var lastNum float64
	var operation string

	for {
		fmt.Println()
		fmt.Println("================================================ \n << WELCOME TO EXCEL'S CALCULATIVE WORLD >> \n================================================")
		fmt.Println()
		fmt.Println("Select operation:")
		fmt.Println("1. +", ": addition")
		fmt.Println("2. -", ": subtraction")
		fmt.Println("3. *", ": multiplication")
		fmt.Println("4. /", ": division")
		fmt.Println("5.  : quit")
		fmt.Println("6.  : help")
		fmt.Println()
		fmt.Scan(&operation)
		fmt.Println()
		if operation == "5" {
			fmt.Println("Goodbye!!!")
			break
		} else if operation == "6" {
			fmt.Println("This are the correct inputs:")
			fmt.Println("+ to add")
			fmt.Println("- to subtract")
			fmt.Println("* to multiply")
			fmt.Println("/ to divide")
		}

		fmt.Print("Enter First Number: ")
		fmt.Scanln(&firstNum)
		fmt.Println()
		fmt.Print("Enter Last Number: ")
		fmt.Scanln(&lastNum)
		fmt.Println()

		if operation == "5" {
			fmt.Println("Goodbye!!!")
			break
		}
		if operation == "6" {
			fmt.Println("This are the correct inputs:")
			fmt.Println("+ to add")
			fmt.Println("- to subtract")
			fmt.Println("* to multiply")
			fmt.Println("/ to divide")
		}

		switch operation {
		case "1":
			fmt.Println("add", firstNum, lastNum)
			fmt.Println()
			fmt.Println("Result:", firstNum+lastNum)
			fmt.Println()
			continue

		case "2":
			fmt.Println("sub", firstNum, lastNum)
			fmt.Println()
			fmt.Println("Result:", firstNum-lastNum)
			fmt.Println()
			continue

		case "3":
			fmt.Println("mul", firstNum, lastNum)
			fmt.Println()
			fmt.Println("Result:", firstNum*lastNum)
			fmt.Println()
			continue

		case "4":
			if lastNum == 0 {
				fmt.Println("Mathematical Error:")
				continue
			} else {
				fmt.Println("div", firstNum, lastNum)
				fmt.Println()
				fmt.Println("Result:", firstNum/lastNum)
				fmt.Println()
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
