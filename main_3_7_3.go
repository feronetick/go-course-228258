package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const Layout = "02.01.2006 15:04:05"

func main() {
	src, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(src)), ",")

	dt1, err := time.Parse(Layout, parts[0])

	if err != nil {
		panic(err)
	}

	dt2, err := time.Parse(Layout, parts[1])

	if err != nil {
		panic(err)
	}

	var dur time.Duration

	if dt1.After(dt2) {
		dur = dt1.Sub(dt2)
	} else {
		dur = dt2.Sub(dt1)
	}

	fmt.Println(dur.Round(time.Second))
}
