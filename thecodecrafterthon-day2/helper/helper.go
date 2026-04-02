package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func BinToDecimal(bin string) int {
	n, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("Not binary input!!")
	}
	return int(n)
}

func HexToDecimal(hex string) int {
	n, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println(err, "is not a valid hex!!")
	}
	return int(n)
}

func DecToBin(dec int64) string {
	n := strconv.FormatInt(dec, 2)
	return strings.ToUpper(n)
}

func DecToHex(dec int64) string {
	n := strconv.FormatInt(dec, 16)
	return strings.ToUpper(n)
}
