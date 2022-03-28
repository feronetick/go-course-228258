package main

import (
	"fmt"
	"os"
)

func scanString() (s string) {
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return s
}

func main() {
	source := scanString()
	result := ""

	for _, v := range source {
		result = string(v) + result
	}

	fmt.Println(result)
}
