package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const defaultCsvPath = "problems.csv"

func main() {
	var csvPath string

	flag.StringVar(&csvPath, "csv", defaultCsvPath, "a csv file in the format of 'question,answer'")
	flag.Parse()

	// file implements io.Reader interface
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// bufio.Reader also implements io.Reader interface
	reader := bufio.NewReader(file)

	// open the CSV file
	csv_reader := csv.NewReader(reader)
	csv_reader.ReuseRecord = true

	idx, score := 0, 0
	for ; ; idx++ {
		records, err := csv_reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		q, ans := records[0], records[1]
		// print the prompt
		fmt.Printf("Problem #%d: %s = ", idx, q)

		// wait for user input
		var input string
		_, err = fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		if ans == input {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, idx)
}
