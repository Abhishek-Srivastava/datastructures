package arrays

import "fmt"

func pairWithSumEquals(a []int, x int) {
	hashMap := make(map[int]int)
	for i, v := range a {
		hashMap[v] = i
	}
	fmt.Println("Hashmap is: ", hashMap)
	for i := 0; i < len(a); i++ {
		fmt.Println("Compliment of ", a[i], " is ", x-a[i])
		if index, ok := hashMap[x-a[i]]; ok {
			fmt.Println("Pair is: ", i, index)
		}
	}
}

/*
// Run se
func Run() {
	pairWithSumEquals([]int{10, 5, 40, 1, 6, 4}, 9)
}
*/
