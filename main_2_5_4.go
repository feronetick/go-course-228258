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
	src := readString()
	out := ""

	for k, v := range src {
		if k%2 != 0 {
			out += string(v)
		}
	}

	fmt.Println(out)
}
