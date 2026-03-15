package main

import (
	"fmt"
)

func PrintIf(str string) string {
	result := "G\n"
	if str == "" {
		return result
	} else if len(str) >= 3 {
		return result
	} else {
		return "Invalid Input"
	}
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
