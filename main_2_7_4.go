package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	out := ""

	for _, s := range ReadString() {
		if n, err := strconv.Atoi(string(s)); err == nil {
			out += fmt.Sprint(n * n)
		} else {
			panic(fmt.Sprintf("invalid input: %s", err))
		}
	}

	fmt.Println(out)
}
