package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	solveRandom()
}

func solve(sl []int32) {
	var wg sync.WaitGroup
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%d\t ", sl[i]*sl[i])
			wg.Done()
		}()
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
}

func solveRandom() {
	rand.Seed(time.Now().UnixNano())
	var (
		amountOfArrays = rand.Int31n(9) + 1
		wg             sync.WaitGroup
	)
	fmt.Printf("Amount of arrays %d\n", amountOfArrays)
	for i := 0; i < int(amountOfArrays); i++ {
		var lengthOfArray = rand.Int31n(9) + 1
		var tmpSlice []int32
		fmt.Printf("\n#Case%d -> ", i+1)
		for j := 0; j < int(lengthOfArray); j++ {
			tmpSlice = append(tmpSlice, rand.Int31n(9)+1)
			fmt.Printf("%d\t ", tmpSlice[j])
		}
		fmt.Printf("\nAnswer -> ")
		solve(tmpSlice)
		fmt.Printf("\n")
		wg.Wait()
	}
}
