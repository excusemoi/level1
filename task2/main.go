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
			fmt.Printf("\t%d ", tmpSlice[j])
		}
		fmt.Printf("\nAnswer -> ")
		for j := 0; j < len(tmpSlice); j++ {
			wg.Add(1)
			go func(num int32) {
				fmt.Printf("\t%d ", num*num)
				wg.Done()
			}(tmpSlice[j])
			time.Sleep(time.Millisecond)
		}
		fmt.Printf("\n")
		wg.Wait()
	}
}
