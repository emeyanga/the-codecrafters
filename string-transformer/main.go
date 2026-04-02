// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Enyo-ojo Meyanga(emeyanga)]
// Squad:  [The Compilers]

package main

import (
	"bufio"
	"fmt"
	"helper/helper"
	"os"
)

var word string

func main() {
	for {
		fmt.Print("Enter the string to transform: ")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		text := scanner.Text()

		lists := []string{"UPPER", "LOWER", "CAP FIRST LETTER", "SNAKE", "TITLE", "REVERSE", "EXIT"}
		for p, race := range lists {
			fmt.Printf("[%d]. %s\n", p+1, race)
		}
		fmt.Print("Select operation: ")
		fmt.Scanln(&word)

		if word == "7" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}
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
			fmt.Println(helper.Title(text))
		case "6":
			fmt.Println(helper.ReverseWord(text))
		default:
			fmt.Println("not a command")

		}
	}
}
