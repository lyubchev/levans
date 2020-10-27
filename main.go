package main

import "fmt"

func main() {

	str1 := "hello0"
	str2 := "helo"
	fmt.Println(lev(str1, str2, len(str1), len(str2)))

}

func lev(str1, str2 string, i, j int) int {
	if min(i, j) == 0 {
		return max(i, j)
	}

	if str1[i-1] == str2[j-1] {
		return lev(str1, str2, i-1, j-1)
	}

	return min(
		min(lev(str1, str2, i-1, j), lev(str1, str2, i, j-1)),
		lev(str1, str2, i-1, j-1),
	) + 1
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
