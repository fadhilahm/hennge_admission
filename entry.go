package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// declare interface for reading standard input
var scanner *bufio.Scanner

// declare slice to hold recursion values until input has been finished
var s []int

// declare new type for showing output
type callbackFunction func(a ...interface{}) (n int, err error)


func main() {
	// assign new `Scanner` type to `scanner`
	scanner = bufio.NewScanner(os.Stdin)

	// use `ScanWords` as the split function for tokenization
	scanner.Split(bufio.ScanWords)

	// read the first token
	if !scanner.Scan() {
		panic("missing number of test cases to follow, main()")
	}

	// convert the first token from string into integers
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("error converting token into integer, main()")
	}

	// execute recursion using `n`
	recursion(n)

	// print every element of slice `s`
	loop(0,len(s), s, fmt.Println)


}

// recursion function will be called equal to the number of case
func recursion(n int) {

	// recursion base case 
	if n == 0 {
		return 
	}

	// read the amount of integers in this current test
	if !scanner.Scan() {
		panic("missing number of integers for the test, recursion()")
	}

	// parse into integer
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("error converting token into integer, recursion()")
	}

	// append the value into the slice
	s = append(s, add(x))

	// execute recursion
	recursion(n - 1)
}

// add function will adds up non-negative integer for a single case
func add(x int) int {

	// recursion base case
	if x == 0 {
		return 0
	}

	// read the next integer
	if !scanner.Scan() {
		panic("missing integer, add()")
	}

	// parse into integer
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("error converting target into integer, add()")
	}

	// turn integer into 0 to handle negative
	if i < 0 {
		i = 0
	}

	// execute recursion
	return i*i + add(x - 1)
}

func loop(index int, maxLen int, i interface{}, cb callbackFunction ) {
	// recursion base case
	if index == maxLen {
		return
	}

	// type switch
	switch i.(type) {
	case []int:
		cb(i.([]int)[index])
	}

	// execute recursion
	loop(index + 1, maxLen, i, cb)
}