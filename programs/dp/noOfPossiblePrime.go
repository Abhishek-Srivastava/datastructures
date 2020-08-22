package dp

import (
	"fmt"
	"strconv"
)

func isPrime(s string, i, j int) bool {
	r := []rune(s)
	n, _ := strconv.Atoi(string(r[i : j+1]))
	fmt.Println("Evaluating", n)
	if n == 1 || n == 0 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var count int

func findMaxNoOfwaystoGeneratePrime(s string, i, j int) int {
	if i > j {
		return 0
	}
	// if t[i][j] != -1 {
	// 	return t[i][j]
	// }
	// if isPrime(s, i, j) {
	// 	return
	// }
	// for k := i; k < j; k++ {
	// 	t[i][j] = findMaxNoOfwaystoGeneratePrime(s, i, k) + findMaxNoOfwaystoGeneratePrime(s, k+1, j)
	// }
	// return count
	if isPrime(s, i, j) {
		fmt.Println(s, "is a prime no")
		count++
	}
	for k := i; k < j; k++ {
		if isPrime(s, i, k) && isPrime(s, k+1, j) {
			fmt.Printf("%s and %s are prime pairs \n", s[i:k+1], s[k+1:j+1])
			count++
		}
	}
	return count
}

func Run() {
	s := "78910"
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
	ans := findMaxNoOfwaystoGeneratePrime(s, 0, n-1)
	fmt.Println("Final answer", ans)
	//fmt.Println(t)
}
