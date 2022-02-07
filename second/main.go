package main

import (
	"fmt"
	"math"
	"math/rand"
)

const numbersCnt = 10

func receiver(ch chan<- float64, n int) {
	for i := 0; i < n; i++ {
		ch <- math.Pow(rand.Float64(), 9)
	}
	close(ch)
}

func writer(ch1 <-chan float64, ch2 chan<- float64) {
	for n := range ch1 {
		ch2 <- n
	}
	close(ch2)
}

func main() {
	fmt.Println("---------- Second ----------")

	ch1 := make(chan float64)
	ch2 := make(chan float64)

	go receiver(ch1, numbersCnt)
	go writer(ch1, ch2)

	for n := range ch2 {
		fmt.Println(n)
	}

	fmt.Println("----------------------------")
}
