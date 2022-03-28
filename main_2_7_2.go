package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	fmt.Println(strings.Join(strings.Split(ReadString(), ""), "*"))
}
