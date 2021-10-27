package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func CountSquaresConc(sl []int64) []int64 {
	var (
		wg  sync.WaitGroup //для создания группы выполняющихся горутин
		res []int64
	)
	for i := 0; i < len(sl); i++ {
		wg.Add(1) //увеличиваем счётчик на 1
		go func(n int64) {
			wg.Done() //уменьшаем счётчик
			res = append(res, n*n)
		}(sl[i])
		time.Sleep(time.Microsecond) /*ожидание завершения горутины (поддержание порядка)
		                               невозможно предугадать порядок выполнение горутин*/
	}
	wg.Wait() // ждём окочания работы горутин
	return res
}

func Solve(sl []int64) {
	res := CountSquaresConc(sl)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d\t", res[i])
	}
}

func SolveRandom() { //произвольное количество массивов с произвольными элементами
	rand.Seed(time.Now().UnixNano())
	var (
		amountOfArrays = rand.Int31n(9) + 1
		wg             sync.WaitGroup
	)
	fmt.Printf("Amount of arrays %d\n", amountOfArrays)
	for i := 0; i < int(amountOfArrays); i++ {
		var lengthOfArray = rand.Int31n(9) + 1
		var tmpSlice []int64
		fmt.Printf("\n#Case%d -> ", i+1)
		for j := 0; j < int(lengthOfArray); j++ {
			tmpSlice = append(tmpSlice, rand.Int63n(9)+1)
			fmt.Printf("%d\t", tmpSlice[j])
		}
		fmt.Printf("\nAnswer -> ")
		Solve(tmpSlice)
		fmt.Printf("\n")
		wg.Wait()
	}
}
