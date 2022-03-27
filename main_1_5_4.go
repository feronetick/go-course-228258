package main

import (
	"fmt"
	"os"
)

func main() {
	var s string

	if _, err := fmt.Scan(&s); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	l := len(s)

	if l > 0 {
		fmt.Println(string(s[l-1]))
	} else {
		fmt.Println(0)
	}
}
