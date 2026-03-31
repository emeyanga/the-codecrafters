package main

import (
	"bufio"
	"fmt"
	"helper/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

var choice, operator, conversion, bin, hex string
var firstNum, lastNum float64
var dec int64
var word string

func main() {

	reader := bufio.NewReader(os.Stdin)

	var input string
	// var num1, num2 float64
	var choice int
	var err error

	for {

		fmt.Print("\n========== FULL-STACK RELOAD =============\n")
		fmt.Println()

		ProgramList := []string{"CALCULATOR", "BASE CONVERTER", "STRING-TRANSFORMER", "EXIT"}
		for i, node := range ProgramList {
			fmt.Printf("[%d]. %s\n\n", i+1, node)
		}
		fmt.Println()
		fmt.Print("Enter program operation: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Print("\n[INPUT ERROR]: ---->> Choice cannot be empty. Try Again\n")
			continue
		}

		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Print("\n[INPUT ERROR]: ---->> Your choice is out of range. Try Again\n")
			continue
		}

		if choice == 4 {
			fmt.Println("Exiting...... Converter!")
			return
		}
		switch choice {
		case 1:
			fmt.Print("\n Time to solve problems!!! \n")
			fmt.Print("\n Enter first number: ")
			fmt.Scanln(&firstNum)
			if firstNum <= 0 && firstNum >= 9 {
				fmt.Println("Not an Integer!!!")
				continue
			}
			// Num, err := strconv.ParseFloat(firstNum, 10, 64)
			// if err != nil {
			// 	fmt.Println("NOT AN INTEGER!!!")
			// 	return
			// }
			// Num = firstNum

			fmt.Print("\n Enter the operator to use:" + "\n" + "1. Add(+)" + "\n" + "2. Subtract(-)" + "\n" + "3. Multiply(*)" + "\n" + "4. Divide(/)" + "\n" + "5. Modulus(%)" + "\n" + "6. Power(^)\n\n")
			fmt.Scanln(&operator)
			fmt.Print("\n Enter second number: ")
			fmt.Scanln(&lastNum)
			switch operator {
			case "+":
				fmt.Println("\n add", firstNum, lastNum)
				fmt.Println("\n Result:", firstNum+lastNum)
				continue
			case "-":
				fmt.Println("\n sub", firstNum, lastNum)
				fmt.Println("\n Result:", firstNum-lastNum)
				continue
			case "*":
				fmt.Println("\n mul", firstNum, lastNum)
				fmt.Println("\n Result:", firstNum*lastNum)
				continue
			case "/":
				if lastNum == 0 {
					fmt.Println("\n Zero division error")
				} else {
					fmt.Println("\n div", firstNum, lastNum)
					fmt.Println("\n Result:", firstNum/lastNum)
					continue
				}
			case "%":
				result := math.Mod(float64(firstNum), float64(lastNum))
				fmt.Println("\n mod", firstNum, lastNum)
				fmt.Println("\n Result:", result)
				continue
			case "^":
				result := math.Pow(float64(firstNum), float64(lastNum))
				fmt.Println("\n pow", firstNum, lastNum)
				fmt.Println("\n", result)
				continue
			}
		case 2:
			fmt.Print("It's time to convert bases!!! \n")
			fmt.Println("Choose conversion: \n 1. binToDecimal(bin) \n 2. hexToDecimal(hex) \n 3. decToBinAndHex(dec)")
			fmt.Scanln(&conversion)

			switch conversion {
			case "1":
				fmt.Println("Input Binary digits:")
				fmt.Scanln(&bin, "\n")
				fmt.Println(helper.BinToDecimal(string(bin)))
				continue
			case "2":
				fmt.Println("Input Hexadecimal digits:")
				fmt.Scanln(&hex, "\n")
				fmt.Println(helper.HexToDecimal(string(hex)))
				continue
			case "3":
				fmt.Println("Input decimal numbers:")
				fmt.Scanln(&dec, "\n\n")
				if dec < 0 && dec > 9 {
					fmt.Println("Input is not a decimal!!")
					continue
				} else {
					fmt.Println("\n", helper.DecToBase(int64(dec)))
				}
			}
		case 3:
			fmt.Print("Enter the string to transform: ")
			scanner := bufio.NewScanner(os.Stdin)

			scanner.Scan()

			text := scanner.Text()

			lists := []string{"UPPER", "LOWER", "CAP FIRST LETTER", "SNAKE", "TITLE", "REVERSE"}
			for p, race := range lists {
				fmt.Printf("[%d]. %s\n", p+1, race)
			}
			fmt.Print("Select operation: ")
			fmt.Scanln(&word)

			switch word {
			case "1":
				fmt.Println(helper.ToUppercase(text))
			case "2":
				fmt.Println(helper.ToLowercase(text))
			case "3":
				fmt.Println(helper.ToCaps(text))
			case "4":
				fmt.Println(helper.ToSnake(text))
			case "5":
				fmt.Println(helper.DoTitle(text))
			case "6":
				fmt.Println(helper.ReverseWord(text))

			default:
				fmt.Println("not a command")

			}

		}
	}
}
