package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

func readString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	text := readString()

	if unicode.IsUpper([]rune(text)[0]) && text[len(text)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
