package goroutines

import (
	"fmt"
	"sync"
)

func printeven(max int, odd chan bool, even chan bool, wg *sync.WaitGroup) {
	num := 0
	for num < max {
		fmt.Println(num)
		even <- true
		num = num + 2
		<-odd
	}
	wg.Done()
}

func printodd(max int, odd chan bool, even chan bool, wg *sync.WaitGroup) {
	num := 1
	for num < max {
		<-even
		fmt.Println(num)
		num = num + 2
		odd <- true
	}
	wg.Done()
}

func Run() {
	even := make(chan bool)
	odd := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	var max int
	max = 100
	go printeven(max, even, odd, &wg)
	go printodd(max, even, odd, &wg)
	wg.Wait()
}
