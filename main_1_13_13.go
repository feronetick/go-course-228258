package main

import (
	"fmt"
	"os"
	"strings"
)

func scanString() (s string) {
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return s
}

func main() {
	fmt.Println(strings.ReplaceAll(scanString(), scanString(), ""))
}
