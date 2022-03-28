package main

import (
	"fmt"
)

func divide(x int, y int) (int, error) {
	// реализовано в тесте
	return 0, nil
}

func main() {
	var a, b int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("ошибка")
		return
	}

	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("ошибка")
		return
	}

	if res, err := divide(a, b); err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(res)
	}
}
