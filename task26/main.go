package task26

import (
	"fmt"
	"strings"
)

func Solve(s string){
	fmt.Printf("Input string: %s\nUnique: %t\n", s, IsStringUnique(s))
}

func IsStringUnique(s string) bool{
	sr := []rune(strings.ToLower(s))
	m := make(map[rune]int)
	for _, r := range sr {
		m[r]++ //map исключает повторение элементов
		if m[r] > 1 {
			return false
		} //если количество элементов больше одного -> ...
	}
	return true
}