package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	Solve(12345, 5)
}

func Solve(num, i int64) {
	fmt.Printf("Number %d -> %b\n", num, num)
	num, err := ReverseBit(num, i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result %d -> %b\n", num, num)
}

func ReverseBit(num, i int64) (int64, error) {
	if i < 0 {
		return num, errors.New("negative shift amount")
	}
	return num ^ 1 << i, nil
}

/*^ - исключающее или
0 ^ 0 = 0
0 ^ 1 = 1
1 ^ 0 = 1
1 ^ 1 = 0
таким образом при выполнении операции
a_1a_2...a_i...a_n = a_1a_2...a_i...a_n << 0_1 0_2...1_i...0_n (a_j = 0 или 1 для любого j)
если i-ый бит равен нулю 0 ^ 1 = 1
если одному 0 ^ 0 = 0
что и требовалось в задании*/
