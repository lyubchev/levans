package main

import (
	"fmt"

	"github.com/impzero/levans"
)

func main() {

	str1 := "hello0"
	str2 := "helo"

	// gives the number of actions to be taken in order to
	// turn str1 into str2
	// output: 2 (remove "0" from str1, remove "l" from str1)
	fmt.Printf("The Levenshtein distance between %s and %s is %d\n", str1, str2, levans.Lev(str1, str2, len(str1), len(str2)))

}
