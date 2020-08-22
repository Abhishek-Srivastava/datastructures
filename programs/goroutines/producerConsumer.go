package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func producer(prod chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		prod <- i
		fmt.Println("Producing ", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func consumer(output <-chan int, done chan<- bool, name string) {
	for {
		output, open := <-output
		if open {
			fmt.Println("Consuming in ", name, output)
		} else {
			fmt.Println("Channel closed")
			done <- true
			return
		}
	}
}

// func Run() {
// 	fmt.Println("Starting producer consumer")
// 	prod := make(chan int, 20)
// 	done := make(chan bool)
// 	go consumer(prod, done, "cosumer1")
// 	wgP := sync.WaitGroup{}
// 	wgP.Add(1)
// 	go producer(prod, &wgP)
// 	wgP.Wait()
// 	close(prod)
// 	// go consumer(prod, done, "consumer2")
// 	<-done
// }
