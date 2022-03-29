package main

import (
	"fmt"
)

func ScanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		panic(fmt.Sprintf("invalid input: %s", err))
	}

	return i
}

func main() {
	fn := func(src uint) uint {
		out := []uint{}

		for i := src; i > 0; i /= 10 {
			n := i % 10

			if n != 0 && n%2 == 0 {
				out = append(out, n)
			}
		}

		res := uint(0)

		for i, v := range out {
			p := uint(1)

			for j := 0; j < i; j++ {
				p *= 10
			}

			res += v * p
		}

		if res == 0 {
			return 100
		}

		return res
	}

	fmt.Println(fn(uint(ScanInt())))
}
