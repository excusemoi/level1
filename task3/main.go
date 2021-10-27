package task3

import (
	"fmt"
	"sync"
)

func Solve(sl []int64) {
	fmt.Println(CountArraySumConc(CountSquaresConc(sl)))
}

//аналогично второму заданию, только теперь time.Sleep ни к чему, ведь нам не важен порядок выполнения горутин

func CountArraySumConc(sl []int64) int64 {
	var sum int64
	var wg sync.WaitGroup

	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func(n int64) {
			wg.Done()
			sum += n
		}(sl[i])
	}
	wg.Wait()
	return sum
}

func CountSquaresConc(sl []int64) []int64 {
	var (
		wg  sync.WaitGroup
		res []int64
	)
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func(n int64) {
			res = append(res, n*n)
			wg.Done()
		}(sl[i])
	}
	wg.Wait()
	return res
}
