package dp

import (
	"math"
)

//Given a string,
// a partitioning of the string is a palindrome partitioning
// if every substring of the partition is a palindrome.
func isPalindrome(s string, i, j int) bool {
	sSlice := []rune(s)
	for i <= j {
		if sSlice[i] != sSlice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

var noOfRecursion int

func solve(s string, i, j int) int {
	if i >= j {
		return 0
	}

	if t[i][j] != -1 {
		return t[i][j]
	}
	if isPalindrome(s, i, j) {
		return 0
	}
	noOfRecursion++
	minVal := math.MaxInt32
	for k := i; k <= j-1; k++ {
		var l, r int
		if t[i][k] == -1 {
			l = solve(s, i, k)
			t[i][k] = l
		} else {
			l = t[i][k]
		}
		if t[k+1][j] == -1 {
			r = solve(s, k+1, j)
			t[k+1][j] = r
		} else {
			r = t[k+1][j]
		}
		temp := 1 + l + r
		if temp < minVal {
			minVal = temp
		}
	}
	t[i][j] = minVal
	return minVal
}

/*
func Run() {
	// s := "nitinxynitin"
	s := "BABABCBADCD"
	n := len(s)
	t = make([][]int, n+1)
	for k := range t {
		t[k] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			t[i][j] = -1
		}
	}
	for k := range t {
		fmt.Println(t[k])
	}
	// intialize the matrix with -1

	fmt.Println(solve(s, 0, len(s)-1))
	fmt.Println("No of recursive calls", noOfRecursion)
}
*/
