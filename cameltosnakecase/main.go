package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	un := []rune(s)
	result := []rune{}

	for i, v := range un {
		if !(v >= 'a' && v <= 'z') && !(v >= 'A' && v <= 'Z') {
			return s
		}
		if i == len(un)-1 && v >= 'A' && v <= 'Z' {
			return s
		}
		if i > 0 && v >= 'A' && v <= 'Z' && un[i-1] >= 'A' && un[i-1] <= 'Z' {
			return s
		}
	}
	for i, v := range un {
		if i > 0 && v >= 'A' && v <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, v)
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
