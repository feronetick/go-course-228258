package main

import "fmt"

func main() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Printf("Invalid input: %s", err)
	} else {
		fmt.Println(a*2 + 100)
	}

}
