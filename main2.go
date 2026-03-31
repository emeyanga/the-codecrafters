package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

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
        
    }
}
