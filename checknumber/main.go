package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	for _, v := range arg {
		if v >= '1' && v <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
