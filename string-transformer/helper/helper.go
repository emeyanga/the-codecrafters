package helper

import (
	"strings"
	"unicode"
)

func ToUppercase(str string) string {
	return strings.Join(strings.Fields(strings.ToUpper(str)), " ")
}
func ToLowercase(str string) string {
	return strings.Join(strings.Fields(strings.ToLower(str)), " ")
}
func ToCaps(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		words[i] = HelpCap(words[i])
	}
	return strings.Join(words, " ")
}
func HelpCap(str string) string {
	title := strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
	return title
}

func WordsCheck(str string) bool {
	smallWords := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

	for i := 0; i < len(smallWords); i++ {
		if strings.Contains(smallWords[i], str) {
			return true
		}
	}
	return false
}
func Title(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		if WordsCheck(words[i]) == true && i > 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = HelpCap(words[i])
		}
	}
	return strings.Join(words, " ")
}

func ToSnake(str string) string {
	var cleaned strings.Builder

	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
			cleaned.WriteRune(char)
		}
	}
	words := strings.ToLower(cleaned.String())

	return strings.Join(strings.Fields(words), "_")
}

func ReverseWord(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		words[i] = SwapChar(words[i])
	}
	return strings.Join(words, " ")
}

func SwapChar(str string) string {
	runes := []rune(str)
	l := len(runes)

	for i := 0; i < l/2; i++ {
		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	}
	return string(runes)
}
