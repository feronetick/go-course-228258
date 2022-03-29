package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	res := 0

	for s.Scan() {
		if n, err := strconv.Atoi(s.Text()); err == nil {
			res += n
		}
	}

	w.WriteString(strconv.Itoa(res))
	w.Flush()
}
