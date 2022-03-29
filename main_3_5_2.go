package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	rootDir := filepath.Join(".", "task_3_5_2")

	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)

		if err != nil {
			return err
		}

		defer f.Close()

		r, err := csv.NewReader(f).ReadAll()

		if err != nil {
			return err
		}

		if len(r) == 10 {
			fmt.Println(path)
			fmt.Println(r[4][2])
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
