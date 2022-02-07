package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const pi = 3.141592653589793

func joinChannels(chs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch <-chan int) {
			for value := range ch {
				out <- value
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func strangeCalculations(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- int(pi*float32(i)) + rand.Intn(13)
		}
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("---------- Third ----------")

	ch1 := strangeCalculations(2)
	ch2 := strangeCalculations(9)
	ch3 := strangeCalculations(11)

	for num := range joinChannels(ch1, ch2, ch3) {
		fmt.Println(num)
	}

	fmt.Println("---------------------------")
}
