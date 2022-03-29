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

	str := strings.TrimSpace(string(src))
	str = strings.ReplaceAll(str, " мин. ", "m")
	str = strings.ReplaceAll(str, " сек.", "s")

	dur, err := time.ParseDuration(str)

	if err != nil {
		panic(err)
	}

	dt := time.Unix(int64(1589570165), 0)

	fmt.Println(dt.Add(dur).Format(time.UnixDate))
}
