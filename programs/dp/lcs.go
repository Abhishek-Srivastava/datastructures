package dp

import (
	"fmt"
	"math"
)

func longestCommonSubsequence(s1, s2 []rune, m, n int) int {
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = int(math.Max(float64(t[i-1][j]), float64(t[i][j-1])))
			}
		}
	}
	return t[m][n]
}

func printLCS(s1, s2 []rune, lenLCS, m, n int) {
	stack := []rune{}
	i := m
	j := n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			stack = append(stack, s1[i-1])
			i--
			j--
		} else {
			if t[i-1][j] > t[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}
	for i := len(stack); i > 0; i-- {
		fmt.Print(string(stack[i-1]))
	}
	fmt.Println()
}

/*
func Run() {
	s1 := []rune{
		'a', 'b', 'c', 'd', 'e',
	}
	s2 := []rune{
		'a', 'e', 'c', 'f',
	}
	m := len(s1)
	n := len(s2)
	t = make([][]int, m+1)
	for i := range t {
		t[i] = make([]int, n+1)
	}
	lenLCS := longestCommonSubsequence(s1, s2, m, n)
	fmt.Println(lenLCS)
	for i := range t {
		fmt.Println(t[i])
	}
	fmt.Println(string(s1))
	fmt.Println(string(s2))
	printLCS(s1, s2, lenLCS, m, n)
}
*/
