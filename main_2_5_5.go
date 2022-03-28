package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	src := readString()
	out := ""

	for _, v := range src {
		s := string(v)

		if strings.Count(src, s) == 1 {
			out += s
		}
	}

	fmt.Println(out)
}
