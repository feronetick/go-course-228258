package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func readString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	text := readString()
	txet := ""

	for _, v := range text {
		txet = string(v) + txet
	}

	if txet == text {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
