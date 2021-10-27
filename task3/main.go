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
	var m sync.Mutex

	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func(n int64) {
			wg.Done()
			m.Lock()
			sum += n
			m.Unlock()
		}(sl[i])
	}
	wg.Wait()
	return sum
}

func CountSquaresConc(sl []int64) []int64 {
	var (
		wg  sync.WaitGroup
		res []int64
		m sync.Mutex
	)
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func(n int64) {
			m.Lock()
			res = append(res, n*n)
			m.Unlock()
			wg.Done()
		}(sl[i])
	}
	wg.Wait()
	return res
}
