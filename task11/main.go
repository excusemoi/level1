package task11

import "fmt"

func Solve(s1, s2 []int64){
	fmt.Printf("First set:\n%v\n", s1)
	fmt.Printf("Second set:\n%v\n", s2)
	fmt.Printf("Result:\n%v\n", Intersection(s1, s2))
}

func Intersection(s1, s2 []int64) (s []int64) {
	m := make(map[int64]struct{})
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				m[s1[i]] = struct{}{}//элементы совпадают - заносим в map, исключая повторение элементов таким образом
			}
		}
	}
	for key, _ := range m {
		s = append(s, key)
	}
	return s
}