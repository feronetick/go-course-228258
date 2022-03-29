package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(".", "task_3_5_3.txt")

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReaderSize(file, 1)
	readed := make([]byte, 1)
	buffer := ""
	index := 1

	for true {
		_, err := reader.Read(readed)

		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			panic(err)
		}

		if string(readed) == ";" {
			if buffer == "0" {
				fmt.Println(index)
				return
			}

			index++
			buffer = ""
		} else {
			buffer += string(readed)
		}
	}
}
