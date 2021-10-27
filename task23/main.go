package task23

import (
	"errors"
	"fmt"
	"log"
)

func Solve(sl []int64, i int64) {
	fmt.Printf("Input array: %v\n", sl)
	if sl_, err := Remove(sl, i); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Printf("Result array: %v\n", sl_)
	}

}

func Remove(sl []int64, i int64) ([]int64, error) {
	if i < 0 || i > int64(len(sl) - 1) {
		return nil, errors.New("Incorrect index")
	}
	return append(sl[:i], sl[i+1:]...), nil
}

//sl[:i] - слайс с элементами [a_1, ..., a_i-1] исходного
//sl[i+1:]... - элементы a_i+1, ..., a_n исходного слайса