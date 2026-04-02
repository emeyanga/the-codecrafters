package main

import (
	"fmt"
	"helper/helper"
)

var conversion, bin, hex string
var dec int64

func main() {
	for {
		fmt.Print("\n ==================================== \n It's time to convert bases!!! \n ==================================== \n")
		fmt.Println("\n Choose conversion: \n 1. binToDecimal(bin) \n 2. hexToDecimal(hex) \n 3. decToBinAndHex(dec) \n 4. Quit")
		fmt.Scanln(&conversion)

		if conversion == "quit" {
			fmt.Println("\n Goodbye, See you soon!!!")
			break
		}
		switch conversion {
		case "1":
			fmt.Println("\n Input Binary digits:")
			fmt.Scanln(&bin, "\n")
			fmt.Println("\n Decimal: ", helper.BinToDecimal(string(bin)))
			continue
		case "2":
			fmt.Println("\n Input Hexadecimal digits:")
			fmt.Scanln(&hex, "\n")
			fmt.Println("\n Decimal: ", helper.HexToDecimal(string(hex)))
			continue
		case "3":
			fmt.Println("\n Input decimal numbers:")
			fmt.Scanln(&dec, "\n")
			if dec < 0 && dec > 9 {
				fmt.Println("\n Input is not a decimal!!")
				continue
			} else {
				fmt.Println()
				fmt.Println("* Binary: ", helper.DecToBin(int64(dec)))
				fmt.Println("* Hex: ", helper.DecToHex(int64(dec)))
				continue
			}
		}
	}
}
