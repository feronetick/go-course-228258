package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Item struct {
	GlobalId       int    `json:"global_id"`
	IsDeleted      int    `json:"is_deleted"`
	Idx            string `json:"Idx"`
	Kod            string `json:"Kod,omitempty"`
	Name           string `json:"Name,omitempty"`
	Razdel         string `json:"Razdel,omitempty"`
	Nomdescr       string `json:"Nomdescr,omitempty"`
	SignatureDate  string `json:"signature_date,omitempty"`
	SystemObjectId string `json:"system_object_id,omitempty"`
}

func main() {
	file, err := os.Open(filepath.Join(".", "task_3_6_2_data.json"))

	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	items := []Item{}

	if err := decoder.Decode(&items); err != nil {
		panic(err)
	}

	sum := 0

	for _, i := range items {
		sum += i.GlobalId
	}

	fmt.Println(sum)
}
