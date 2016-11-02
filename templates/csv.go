package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var example = `
"hello world",nothing to be feared

after blank line, this is followed after a whitespace
`

func main() {
	readSingle()
	readAll()
}

func readSingle() {
	fmt.Println("Read Single line")

	r := csv.NewReader(strings.NewReader(example))

	record, err := r.Read()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("each value:")
	for _, v := range record {
		fmt.Printf("\t\"%v\"\n", v)
	}

}

func readAll() {
	fmt.Println("Read All at once!")

	r := csv.NewReader(strings.NewReader(example))

	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, record := range records {
		fmt.Printf("record: %v\n", record)
		fmt.Println("each value:")
		for _, v := range record {
			fmt.Printf("\t\"%v\"\n", v)
		}
	}

}
