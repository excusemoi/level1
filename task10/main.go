package task10

import (
	"fmt"
	"math"
)

func mapTemperature(sl []float64) (m map[int64][]float64) {
	m = make(map[int64][]float64)
	for i := 0; i < len(sl); i++ {
		rounded := int64(math.Ceil(sl[i]))
		k := rounded - rounded % 10
		m[k] = append(m[k], sl[i])//группируем элементы по их второму разряду
	}
	return m
}

func Solve(sl []float64) {
	fmt.Println("Input array:")
	for i := 0; i < len(sl); i++ {
		fmt.Print(sl[i], " ")
	}
	fmt.Printf("\nResult:\n")
	res := mapTemperature(sl)
	for key, val := range res {
		fmt.Printf("%d: \t%v\n", key, val)
	}
}