package arrays

import "fmt"

var Arr1 []int
var Arr2 []int

func EqualArrays() {
	Arr1 = []int{1, 2, 3, 4, 5, 1, 1, 2, 3}
	Arr2 = []int{2, 1, 4, 3, 5, 2, 1, 1, 3}

	if len(Arr1) != len(Arr2) {
		fmt.Println("Not equal, unequal length")
		return
	}
	myHash := make(map[int]int)

	for _, item := range Arr1 {
		reps, ok := myHash[item]
		if !ok {
			myHash[item] = 1
		} else {
			myHash[item] = reps + 1
		}
	}
	fmt.Println("My hash has: ", myHash)
	for _, item := range Arr2 {
		reps, ok := myHash[item]
		if !ok {
			fmt.Println("Not equal")
			return
		}
		myHash[item] = reps - 1
	}

	for k, _ := range myHash {
		if myHash[k] != 0 {
			fmt.Println("Not equal")
			return
		}
	}
	fmt.Println("equal arrays")
}

/*
func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{}

	for i := 0; i < len(a1); i++ {
		prod := 1
		for j := 0; j < len(a1); j++ {
			if j != i {
				prod = prod * a1[j]
			}
		}
		a2 = append(a2, prod)
	}
	fmt.Println(a1)

	prefixProd := make([]int, 5)
	suffixProd := make([]int, 5)

	prefixProd[0] = a1[0]
	for i := 1; i < len(a1); i++ {
		prefixProd[i] = prefixProd[i-1] * a1[i-1]
	}
	suffixProd[len(a1)-1] = 1
	for i := len(a1) - 2; i >= 0; i-- {
		suffixProd[i] = suffixProd[i+1] * a1[i+1]
	}
	fmt.Println(prefixProd)
	fmt.Println(suffixProd)
	a3 := []int{}
	for i := 0; i < len(a1); i++ {
		a3 = append(a3, prefixProd[i]*suffixProd[i])
	}
	fmt.Println(a2)
	fmt.Println(a3)
}


[1 2 3 4 5]
2*3*4*5
1*3*4*5
1*2*4*5
1*2*3*5
1*2*3*4
*/
