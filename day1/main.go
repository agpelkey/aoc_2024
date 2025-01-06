package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func split_list(s string) []string {
	result := strings.Split(s, "\n")

	fmt.Println(result)
	return result
}

//go:embed input.txt
var input string

//func least_to_greatest()

//func find_the_difference()

//func add()

func main() {

	// Get the first column from the list
	split_list(input)

	// Get the second column from the list

	// Order first list from least to greatest

	// Order second list from least to greatest

	// iterate through the two lists and find the difference between the numbers

	// Add the difference together to get the final answer
}
