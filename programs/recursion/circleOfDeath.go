package recursion

import "fmt"

func deleteIndex(index int, array []int) []int {
	fmt.Println("Deleting: ", array[index], "from", array)
	copy(array[index:], array[index+1:])
	array = array[:len(array)-1]
	return array
}

func circleOfDeath(a []int, killPerson, nextPerson int) int {
	if len(a) == 1 {
		return a[0]
	}
	deletePos := (nextPerson + killPerson - 1) % len(a)
	na := deleteIndex(deletePos, a)
	return circleOfDeath(na, killPerson, deletePos)
}

func Run() {
	a := []int{}
	for i := 1; i <= 40; i++ {
		a = append(a, i)
	}
	fmt.Println(circleOfDeath(a, 7, 0))
}
