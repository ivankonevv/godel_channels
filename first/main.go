package main

import (
	"fmt"
	"math/rand"
)

const numbersCnt = 10

func randNumsGenerator(n int) (ch <-chan int) {
	channel := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			channel <- rand.Int()
		}
		close(channel)
	}()
	return channel
}

func main() {
	fmt.Println("---------- First ----------")

	for num := range randNumsGenerator(numbersCnt) {
		fmt.Println(num)
	}

	fmt.Println("---------------------------")
}
