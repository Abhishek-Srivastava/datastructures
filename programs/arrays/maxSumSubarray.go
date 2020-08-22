package arrays

/*
kadane's algorithm
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func kadanesAlgo(a []int) int {
	curSum := a[0]
	globalSum := a[0]
	for i := 1; i < len(a); i++ {
		curSum = max(a[i], a[i]+curSum)
		if curSum > globalSum {
			globalSum = curSum
		}
	}
	return globalSum
}

/*
// Run se
func Run() {
	a := []int{10, -1, -2, 5, 7, -3}
	fmt.Println(kadanesAlgo(a))
}
*/

//  [4, 5, 6, 7, 8, 9, 1, 2, 3]
//  [7 ,8, 9, 1, 2, 3, 4, 5, 6]
