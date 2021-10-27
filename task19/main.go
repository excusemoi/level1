package task19

import "fmt"

func Solve(s string){
	fmt.Printf("Input string:\n%s\n", s)
	fmt.Printf("Result:\n%s\n", ReverseString(s))
}

func ReverseString(s string) string {
	var s_ string
	var sr  = []rune(s)
	for i, _ := range sr {
		s_ += string(sr[len(sr)-i-1])
	}
	return s_
}

//len = 3
//i = 0 i_ = 3-0-1=2
//i = 1 i_ = 3-1-1=1
//...