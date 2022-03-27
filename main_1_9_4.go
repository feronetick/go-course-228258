package main

import (
	"fmt"
	"os"
)

func scanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return i
}

func main() {
	a := scanInt()

	if a%400 == 0 || (a%4 == 0 && a%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
