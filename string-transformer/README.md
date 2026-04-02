# String Transformer (Go CLI Tool)

## What This Program Does

The String Transformer is a simple command-line program written in Go that allows you to transform text in different ways.

You enter a sentence, choose an operation, and the program outputs the transformed version.

### Available Operations:

1. Convert text to UPPERCASE
2. Convert text to lowercase
3. Capitalize first letter of each word
4. Convert text to snake_case
5. Convert text to Title Case (ignores small words like "and", "the", etc.)
6. Reverse each word
7. Exit the program

### Run the program

From the base directory, run:

```bash
go run main.go
```
## Examples:

### 1. UPPERCASE
Enter the string to transform: hello world
Select operation: 1
HELLO WORLD

### 2. lowercase
Enter the string to transform: HELLO WORLD
Select operation: 2
hello world

### 3. Capitalize First Letter
Enterhe string to transform: hello world
Select operation: 3
Hello World

### 4. snake_case
Enter the string to transform: Hello World Go Lang
Select operation: 4
hello_world_go_lang

### 5. Title Case
Enter the string to transform: the lord of the rings
Select operation: 5
The Lord of the Rings

### 6. Reverse Each Word
Enter the string to transform: hello world
Select operation: 6
olleh dlrow

### 7. Exit
Select operation: 7
Shutting down String Transformer. Goodbye.

## Note:

* The program runs in a loop until you choose **Exit (7)**.
* It removes extra spaces automatically.
* Special characters are removed in snake_case conversion.
* Each word is handled separately for most transformations.