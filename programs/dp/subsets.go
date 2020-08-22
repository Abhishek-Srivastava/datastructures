package dp

// if there's a subset who sum is a given no

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Itob(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func find(a []int, n, sum int) int {
	if n == 0 && sum == 0 {
		return 1
	}
	if n == 0 && sum > 0 {
		return 0
	}
	if t[n][sum] != -1 {
		return t[n][sum]
	}
	if a[n-1] <= sum {
		t[n][sum] = Btoi(Itob(find(a, n-1, sum-a[n-1])) || Itob(find(a, n-1, sum)))
	} else {
		t[n][sum] = find(a, n-1, sum)
	}
	return t[n][sum]
}

func findIterative(a []int, n, sum int) int {
	for i := 1; i < n+1; i++ {
		for j := range t[i] {
			if a[i-1] <= j {
				//t[i][j] = Btoi(Itob(t[i-1][j-a[i-1]]) || Itob(t[i-1][j]))
				t[i][j] = t[i-1][j-a[i-1]] + t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][sum]
}

// minSumDiff
func minSumDiff(a []int, n, sum, minSum int) int {
	for i := 1; i < n+1; i++ {
		for j := range t[i] {
			if a[i-1] <= j {
				//t[i][j] = Btoi(Itob(t[i-1][j-a[i-1]]) || Itob(t[i-1][j]))
				t[i][j] = t[i-1][j-a[i-1]] + t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][sum]
}

var t [][]int

/*
func Run() {
	a := []int{2, 4, 6, 8, 9, 32}
	//sum := 10
	totalSum := 0
	for _, item := range a {
		totalSum += item
	}
	fmt.Println("Total sum", totalSum)
	rows := len(a) + 1
	columns := totalSum + 1
	t = make([][]int, rows)
	for i := range t {
		t[i] = make([]int, columns)
		for j := range t[i] {
			t[i][j] = -1
			if i == 0 {
				t[i][j] = 0
			}
			if j == 0 {
				t[i][j] = 1
			}
		}
	}
	fmt.Println(findIterative(a, len(a), totalSum))
	hashMap := make(map[int]int)
	//lastrow := t[len(a)]
	for possibleSum, items := range t[len(a)] {
		if items > 0 {
			// if there's a sum possible
			hashMap[possibleSum] = items
		}
	}
	fmt.Println(hashMap)
	//fmt.Println(lastrow)
}
*/
