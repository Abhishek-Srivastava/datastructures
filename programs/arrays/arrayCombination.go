package arrays

import (
	"fmt"
	"strconv"
)

func isValid(s string) bool {
	intVal, _ := strconv.Atoi(s)
	if intVal > 0 && intVal <= 26 {
		return true
	}
	return false
}

func totalWays(a string, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	var i int
	for i = 2; i < n+1; i++ {

		if a[i-1] > '0' {
			dp[i] = dp[i-1]
		}
		if a[i-2] == '1' || (a[i-2] == '2' && a[i-1] < '7') {
			dp[i] += dp[i-2]
		}
	}
	fmt.Println(dp)
	return dp[n]
}

/*
func Run() {
	codedPass := "12323"
	fmt.Println(totalWays(codedPass, len(codedPass)))
}*/
