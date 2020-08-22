package sorting

import "fmt"

func partition(a []int, l, h int) int {
	p := l
	i := l
	j := h
	fmt.Println("i, j:", i, j)
	for i < j {
		for {
			i++
			if i == h || a[i] >= a[p] {
				break
			}
		}
		for {
			j--
			if j == l || a[j] <= a[p] {
				break
			}
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[p], a[j] = a[j], a[p]
	fmt.Println(a)
	return j
}

// Quicksort se
func Quicksort(a []int, l, h int) {
	if l < h {
		j := partition(a, l, h)
		Quicksort(a, l, j)
		Quicksort(a, j+1, h)
	}
}

// Run se
/*
func Run() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	Quicksort(a, 0, len(a))
	fmt.Println(a)
}
*/
