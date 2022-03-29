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

type Battery struct {
	charge int
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.charge))
}

func main() {
	batteryForTest := Battery{strings.Count(ReadString(), "1")}

	fmt.Println(batteryForTest)
}
