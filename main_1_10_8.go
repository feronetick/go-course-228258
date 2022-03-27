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
	a := scanString()
	b := scanString()

	res := make([]string, 0)

	for _, r := range a {
		if strings.ContainsRune(b, r) {
			res = append(res, string(r))
		}
	}

	fmt.Println(strings.Join(res, " "))
}
