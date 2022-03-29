package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const Layout = "2006-01-02 15:04:00"

func main() {
	src, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	dt, err := time.Parse(Layout, strings.TrimSpace(string(src)))

	if err != nil {
		panic(err)
	}

	if dt.Hour() >= 13 {
		dt = dt.AddDate(0, 0, 1)
	}

	fmt.Println(dt.Format(Layout))
}
