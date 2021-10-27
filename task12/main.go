package task12

import "fmt"

func Solve(s []string){
	fmt.Printf("Input string sequence:\n%v\n", s)
	fmt.Printf("Result string set:\n%v\n", CreateSet(s))
}

func CreateSet(s []string) (res []string) {
	m := make(map[string]struct{})
	for _, val := range s {
		m[val] = struct{}{}
	}
	for key, _ := range m {
		res = append(res, key)
	}
	return res
}