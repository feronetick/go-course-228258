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

func containsDuplicates(s string) bool {
	for v, r := range s {
		if strings.ContainsRune(s[v+1:], r) {
			return true
		}
	}

	return false
}

func main() {
	if containsDuplicates(scanString()) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
