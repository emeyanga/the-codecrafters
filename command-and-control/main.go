package main

import (
	"fmt"
	"math"
	"strconv"
)

var operation string
var firstNum, lastNum float64
var operator string
var conversion string
var bin, hex string
var dec int64

func binToDecimal(bin string) int {
	n, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("Not binary input!!")
	}
	return int(n)
}

func hexToDecimal(hex string) int {
	n, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println(err, "is not a valid hex!!")
	}
	return int(n)
}

func decToBase(dec int64) string {
	n := strconv.FormatInt(dec, 2) + "\n" + strconv.FormatInt(dec, 16)
	return n
}

func main() {
	for {
		fmt.Println()
		fmt.Print("WELCOME!!!\n\n")
		fmt.Print("Select the function you want to begin:" + "\n" + "1. Calculate (calc)" + "\n" + "2. Base Conversion (base)" + "\n" + "3. String Transform (trans)\n\n")
		fmt.Scanln(&operation)

		switch operation {
		case "1":
			fmt.Print("\n Time to solve problems!!! \n")
			fmt.Print("\n Enter first number: ")
			fmt.Scanln(&firstNum)
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
		case "2":
			fmt.Print("It's time to convert bases!!! \n")
			fmt.Println("Choose conversion: \n 1. binToDecimal(bin) \n 2. hexToDecimal(hex) \n 3. decToBinAndHex(dec)")
			fmt.Scanln(&conversion)
			switch conversion {
			case "1":
				fmt.Println("Input Binary digits:")
				fmt.Scanln(&bin, "\n")
				fmt.Println(binToDecimal(string(bin)))
				continue
			case "2":
				fmt.Println("Input Hexadecimal digits:")
				fmt.Scanln(&hex, "\n")
				fmt.Println(hexToDecimal(string(hex)))
				continue
			case "3":
				fmt.Println("Input decimal numbers:")
				fmt.Scanln(&dec, "\n\n")
				if dec < 0 && dec > 9 {
					fmt.Println("Input is not a decimal!!")
					continue
				} else {
					fmt.Println("\n",decToBase(int64(dec)))
				}
			}
		}
	}
}
