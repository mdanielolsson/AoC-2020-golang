package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := readValues(file)

	var result int
	for _, value := range values {
		found := findOneAddingUp(&values, value)
		if found == 0 {
			continue
		}
		result = found * value
		break
	}

	fmt.Printf("Answer to part 2: %v\n", result)

	for _, value := range values {
		found := findTwoAddingUp(&values, value)
		if found == nil {
			continue
		}
		result = multiplyAll(found)
		break
	}

	fmt.Printf("Answer to part 2: %v\n", result)
}

// multiplyAll will multiply all values of a slice of integers and return the result
func multiplyAll(items []int) int {
	result := 1
	for _, v := range items {
		result = result * v
	}
	return result
}

// Find one that adds up to 2020.
// Takes one int and
func findOneAddingUp(values *[]int, i int) int {
	for _, v := range *values {
		if (v + i) == 2020 {
			return v
		}
	}
	return 0
}

func findTwoAddingUp(values *[]int, i int) []int {
	var result []int
	for _, v := range *values {
		test := findOneAddingUp(values, v+i)
		if test == 0 {
			continue
		}
		result = append(result, v, i, test)
		break
	}

	return result
}

// readValues takes input file descriptor and populate and return slice of all values in the file
func readValues(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	var values []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return values
}
