package arrays

func binSearch(a []int, l, h, x int) int {
	if h-l == 1 {
		if a[l] == x {
			return l
		}
		return -1
	}
	m := (l + h) / 2
	if x == a[m] {
		return m
	}
	if x < a[m] {
		return binSearch(a, l, m, x)
	} else {
		return binSearch(a, m+1, h, x)
	}
}

/*
// modifiedBinSearch for rotated array
func modifiedBinSearch(a []int, x int) int {
	mid := len(a) / 2
	if x == a[mid] {
		return mid
	}
	//  70, 10, 20, 30, 40, 50, 60
	//  60, 70, 10, 20, 30, 40, 50
	//  50, 60, 70, 10, 20, 30, 40
	//  40, 50, 60, 70, 10, 20, 30
	// left side is sorted
	if a[0] < a[mid] {
		return 0
	}
	// right side is sorted
	//if a[len(a)-1] > a[mid]
}*/

// func Run() {
// 	a := []int{10, 20, 30, 40, 50, 60, 70}
// 	fmt.Println(binSearch(a, 0, len(a), 14))
// }
