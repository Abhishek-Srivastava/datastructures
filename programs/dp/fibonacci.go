package dp

// 1 1 2 3 5 8 13 21 34 55

var myMar map[int]int

func fibInefficient(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibInefficient(n-1) + fibInefficient(n-2)
}

func fib(n int) int {
	_, ok := myMar[n]
	if !ok {
		myMar[n] = fib(n-1) + fib(n-2)
		//myMar[n] = myMar[n-1] + myMar[n-2]
	}
	return myMar[n]
}

// func Run() {
// 	myMar = make(map[int]int)
// 	myMar[1] = 1
// 	myMar[2] = 1
// 	fmt.Println(fib(50))
// 	sum := 0
// 	for _, val := range myMar {
// 		sum += val
// 	}
// 	fmt.Println(sum)
// 	// fmt.Println(fibInefficient(50))
// }
