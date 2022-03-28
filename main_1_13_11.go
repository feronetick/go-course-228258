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
	var (
		n = scanInt()
		a = 1 // не с ноля с хуя-то))
		b = 1
		i = 1
	)

	for true {
		i++
		a, b = b, a+b

		if a == n {
			fmt.Println(i)
			return
		} else if a > n {
			fmt.Println(-1)
			return
		}
	}
}
