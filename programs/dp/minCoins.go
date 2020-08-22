package dp

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countNoOfWays(a []int, n, sum int) int {
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if a[i-1] <= j {
				t[i][j] = min(t[i][j-a[i-1]]+1, t[i-1][j])
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][sum]
}

/*
func Run() {
	a := []int{1, 2, 3}
	//sum := 10
	totalSum := 5
	// for _, item := range a {
	// 	totalSum += item
	// }
	fmt.Println("Total sum", totalSum)
	rows := len(a) + 1
	columns := totalSum + 1
	t = make([][]int, rows)
	for i := range t {
		t[i] = make([]int, columns)
		for j := range t[i] {
			t[i][j] = -1
			if i == 0 {
				t[i][j] = math.MaxInt16 - 1
			}
			if j == 0 {
				t[i][j] = 0
			}
		}
	}
	// for i := 1; i < sum + 1; i++ {
	// 	if i % a[0] == 0  {
	// 		t[1][i] =
	// 	}
	// }

	fmt.Println(countNoOfWays(a, len(a), totalSum))
	for i := range t {
		fmt.Println(t[i])
	}
}*/
