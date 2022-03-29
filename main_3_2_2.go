package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func atoi64(src string) int64 {
	out := []rune{}

	for _, r := range src {
		if unicode.IsDigit(r) {
			out = append(out, r)
		}
	}

	res, err := strconv.ParseInt(string(out), 10, 64)

	if err != nil {
		panic(fmt.Sprintf("invalid input: %s", err))
	}

	return res
}

func adding(a string, b string) int64 {
	return atoi64(a) + atoi64(b)
}
