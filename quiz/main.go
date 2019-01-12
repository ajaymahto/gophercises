package main

import (
	// All imported packages must be in alphabetic order.
	// Required for reading a csv file.
	"encoding/csv"
	// Required for standard input/output.
	"fmt"
	// Required for reading commandline flags.
	"flag"
	// Required for trimming spaces.
	"strings"
	// Required for reading a file.
	"os"
)

func main() {
	// flag format -- name, default_val, desc
	csvReader := flag.String("csv", "problems.csv", "CSV file which contains problems in 'ques,ans' format.")

	// To parse all flags defined above.
	flag.Parse()

	// *csvReader is used as the flag package stores the value as pointer.
	file, err := os.Open(*csvReader)
	if err != nil {
		// Custom exit functon to print message and then quit.
		exit(fmt.Sprintf("Error reading file %s : %s", *csvReader, err))
	}

	// New Reader to read csv file. MOst readers implement io.Reader.
	r := csv.NewReader(file)
	// ReadAll() returns a 2-D slice with lines in the csv file as a slice.
	data, err := r.ReadAll()
	if err != nil {
		exit("Error fetching questions!")
	}

	// Custom function to convert a 2-D slice into slice of structs.
	problems := GetProblems(data)

	// To store answer.
	var a string

	// To store scores.
	ctr := 0
	for i, prob := range(problems) {
		fmt.Printf("#%d. %s = ", i+1, prob.q)

		// Scanf is similar to C's scanf and takes address of the variable.
		fmt.Scanf("%s\n", &a)
		if prob.a == a {
			ctr++
		}
	}
	fmt.Printf("Result: %d out of %d\n", ctr, len(problems))
}

// Struct is similar to that in C.
type problem struct {
	q string
	a string
}

// Function to convert the 2-D slice into slice of structs.
func GetProblems(data [][]string) []problem {
	ret := make([]problem, len(data))
	for i, n := range(data) {
		ret[i] = problem{

			// TrimSpace to trim spaces in the retrieved csv file fields.
			q: strings.TrimSpace(n[0]),
			a: strings.TrimSpace(n[1]),
		}
	}
	return ret
}

// Custom exit function.
func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
