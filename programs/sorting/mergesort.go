package sorting

import (
	"fmt"
)

func merge(a1 []int, a2 []int) []int {
	temp := make([]int, len(a1)+len(a2))
	var i, j, k int
	for {
		if a1[i] <= a2[j] {
			temp[k] = a1[i]
			i++
		} else {
			temp[k] = a2[j]
			j++
		}
		k++
		if i == len(a1) || j == len(a2) {
			break
		}
	}

	for i < len(a1) {
		temp[k] = a1[i]
		i++
		k++
	}

	for j < len(a2) {
		temp[k] = a2[j]
		j++
		k++
	}

	return temp
}

func mergeSort(a []int, l, h int) []int {
	if h-l == 1 {
		return a[l:h]
	}
	m := (l + h) / 2
	a1 := mergeSort(a, l, m)
	a2 := mergeSort(a, m, h)
	fmt.Println("merging: ", a1, a2)
	return merge(a1, a2)
}

/*
[1,10, 5, 2, 11, 7]

a[0],a[1]
a[2], a[3]
a[4], a[5]

a[0:2], a[2:4]

*/

/*
func twoWayMerge(a []int) {
	l := float64(len(a))
	for c := 1; c < int(math.Log2(l)); c++ {
		for i := 0; i < len(a); i++ {
			var l, m, h int
			merge(a[l:m], a[m:h])
			a[i:i+c], a[i+c, i+c+c]



		}
	}
}*/

//Run se
func Run() {
	a := []int{10, 5, 3, 40, 20, 2}
	temp := mergeSort(a, 0, len(a))
	fmt.Println(temp)
}
