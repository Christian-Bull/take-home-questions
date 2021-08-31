package main

import (
	"fmt"
	"strings"
)

func main() {

	// #1
	listOne := []int{1, 5, 2, 1, 10}
	printInt(listOne, 6)

	// #2
	s := "MyString"
	fmt.Println(rotateChars(s, 2))
}

// #1 Print the number of integers in an array that are above the given input
// and the number that are below, e.g. for the array [1, 5, 2, 1, 10]
// ex. with input 6, print “above: 1, below: 4”.

// printInt outputs the number of integers that are above and below a given input
// given a list of integers
func printInt(l []int, input int) {
	var numAbove int
	var numBelow int

	for i := 0; i < len(l); i++ {
		if l[i] > input {
			numAbove += 1
		} else if l[i] < input {
			numBelow += 1
		}
	}

	fmt.Printf("above: %d, below: %d\n", numAbove, numBelow)
}

// #2 Rotate the characters in a string by a given input and have the overflow
// appear at the beginning, e.g. “MyString” rotated by 2 is “ngMyStri”.

// rotateChars takes a string + input then returns a rotated string such that
// the overflow chars appear at the beginning
func rotateChars(s string, input int) string {

	a := strings.Split(s, "")

	overflow := strings.Join(a[len(a)-input:], "")
	sMinusOverflow := strings.Join(a[:len(a)-input], "")

	return overflow + sMinusOverflow
}

// #3 If you could change 1 thing about your favorite
// framework/language/platform (pick one), what would it be?
//
// Python:
// Better native support for multithreading/concurency. Currently there's a few
// libraries that achieve this such as concurrent.futures, but they have their
// limitations and often don't offer much of a performance improvement.
