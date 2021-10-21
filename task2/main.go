package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func CountSquaresConc(sl []int32) []int32 {
	var (
		wg  sync.WaitGroup
		res []int32
	)
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func() {
			res = append(res, sl[i]*sl[i])
			wg.Done()
		}()
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
	return res
}

func Solve(sl []int32) {
	res := CountSquaresConc(sl)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d\t", res[i])
	}
}

func SolveRandom() {
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
			fmt.Printf("%d\t", tmpSlice[j])
		}
		fmt.Printf("\nAnswer -> ")
		Solve(tmpSlice)
		fmt.Printf("\n")
		wg.Wait()
	}
}
