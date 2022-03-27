package main

import (
	"fmt"
	"os"
	"strconv"
)

func scanString() (s string) {
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return s
}

func sum(value string) (res int) {
	for _, v := range value {
		if i, err := strconv.Atoi(string(v)); err != nil {
			fmt.Printf("Invalid input: %s", err)
			os.Exit(1)
		} else {
			res += i
		}
	}

	return res
}

func main() {
	s := scanString()

	if sum(s[0:3]) == sum(s[3:6]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
