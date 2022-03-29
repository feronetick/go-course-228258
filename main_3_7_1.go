package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	src, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	dt, err := time.Parse(time.RFC3339, strings.TrimSpace(string(src)))

	if err != nil {
		panic(err)
	}

	fmt.Println(dt.Format(time.UnixDate))
}
