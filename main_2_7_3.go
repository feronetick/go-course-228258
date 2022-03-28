package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() (res string) {
	res, _ = reader.ReadString('\n')
	return res
}

func main() {
	numbers := strings.Split(ReadString(), "")
	sort.Sort(sort.Reverse(sort.StringSlice(numbers)))
	fmt.Println(numbers[0])
}
