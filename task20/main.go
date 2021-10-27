package task20

import (
	"fmt"
	"strings"
)

func Solve(s string) {
	fmt.Printf("Input string:\n%s\n", s)
	fmt.Printf("Result:\n%s\n", ReverseWordsInString(s))
}

//по аналогии с прошлым заданием

func ReverseWordsInString(s string) string {
	arr := strings.Split(s, " ")
	var arr_ []string
	for i, _ :=  range arr {
		arr_ = append(arr_, arr[len(arr) - i - 1] )
	}
	return strings.Join(arr_, " ")
}