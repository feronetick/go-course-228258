package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func ParseFloat(src string) float64 {
	out := []rune{}

	for _, r := range src {
		if unicode.IsDigit(r) || r == '.' {
			out = append(out, r)
		} else if r == ',' {
			out = append(out, '.')
		}
	}

	res, err := strconv.ParseFloat(string(out), 64)

	if err != nil {
		panic(fmt.Sprintf("invalid input: %s", err))
	}

	return res
}

func main() {
	reader := csv.NewReader(bufio.NewReader(os.Stdin))
	reader.Comma = ';'

	row, err := reader.Read()

	if err != nil {
		panic(fmt.Sprintf("invalid input: %s", err))
	}

	if len(row) < 2 {
		panic("invalid input")
	}

	fmt.Printf("%.4f", ParseFloat(row[0])/ParseFloat(row[1]))
}
