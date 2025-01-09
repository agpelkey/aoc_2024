package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func split_list(lines string) ([]int, []int) {

	//fmt.Println(result)
	//return result
	list1 := []int{}
	list2 := []int{}
	for _, line := range lines {
		result := strings.Split(lines, "\n")
		//result := strings.Split(line, "\n")
		//fmt.Println(line)
		first_num, err := strconv.Atoi(result[0])
		if err != nil {
			fmt.Println(err)
		}
		second_num, err := strconv.Atoi(result[1])
		if err != nil {
			fmt.Println(err)
		}
		list1 = append(list1, first_num)
		list2 = append(list2, second_num)
	}

	return list1, list2
}

//go:embed input.txt
var input string

//func least_to_greatest()

//func find_the_difference()

//func add()

func main() {

	// Get the first column from the list
	fmt.Println(split_list(input))

	// Get the second column from the list

	// Order first list from least to greatest

	// Order second list from least to greatest

	// iterate through the two lists and find the difference between the numbers

	// Add the difference together to get the final answer
}
