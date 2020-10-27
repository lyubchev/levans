package levans

func Lev(str1, str2 string, i, j int) int {
	// eventually one of the variables - i or j will be 0 because
	// of the recursion and the other variable will stay
	// unchanged and we have to return it since we keep
	// track of the amount of edits we've performed
	if min(i, j) == 0 {
		return max(i, j)
	}

	// if the last characters of the string match, we ignore them
	if str1[i-1] == str2[j-1] {
		return Lev(str1, str2, i-1, j-1)
	}

	// we calculcate which of the 3 actions - inserting, replacing and deleting
	// has the lowest cost and then we return it by adding 1 (cause we performed one more action)
	return min(
		min(Lev(str1, str2, i-1, j), Lev(str1, str2, i, j-1)),
		Lev(str1, str2, i-1, j-1),
	) + 1
}

// Wroto these helper funcs because I don't want to cast when using go's built-in math lib

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
