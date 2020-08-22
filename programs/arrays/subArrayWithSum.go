package arrays

import (
	"fmt"
)

func findSubArrayWithSum(a []int, x int) {
	var l, r, curSum int
	for {
		if r == len(a) && r == len(a) {
			break
		}
		if curSum < x {
			curSum += a[r]
			r++
		} else if curSum > x {
			curSum -= a[l]
			l++
		}
		fmt.Println("Cursum: ", curSum)
		if curSum == x {
			break
		}
	}
	if curSum != x {
		fmt.Println("Couldn't find a valid subArry with sum: ", x)
	} else {
		fmt.Println("Index:", l, r, "Sub array: ", a[l:r])
	}
}

/*
func Run() {
	a := []int{1, 4, 20, 3, 10, 5}

	// sort.Slice(a, func(i, j int) bool {
	// 	if a[i] <= a[j] {
	// 		return true
	// 	}
	// 	return false
	// })

	sum := 13
	fmt.Println("Input Array: ", a, "X: ", sum)
	findSubArrayWithSum(a, sum)
}
*/
