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
	n := scanInt()
	c := scanInt()
	d := scanInt()

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
