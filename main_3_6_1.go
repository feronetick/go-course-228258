package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Result struct {
	Average float64
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	group := new(Group)

	if err := json.Unmarshal(data, &group); err != nil {
		panic(err)
	}

	count := 0

	for _, student := range group.Students {
		count += len(student.Rating)
	}

	var average float64 = float64(count) / float64(len(group.Students))

	if res, err := json.MarshalIndent(Result{average}, "", "    "); err == nil {
		fmt.Println(string(res))
	} else {
		panic(err)
	}
}
