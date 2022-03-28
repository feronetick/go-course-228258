package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

var reader = bufio.NewReader(os.Stdin)

func readString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func fail() {
	fmt.Println("Wrong password")
	os.Exit(0)
}

func main() {
	src := readString()

	if utf8.RuneCountInString(src) < 5 {
		fail()
	}

	for _, v := range src {
		if !unicode.IsDigit(v) && !unicode.Is(unicode.Latin, v) {
			fail()
		}
	}

	fmt.Println("Ok")
}
