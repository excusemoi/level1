package task13

import "fmt"

func Solve(n1, n2 int64) {
	fmt.Printf("Input numbers:\nn1: %d, n2: %d\n", n1, n2)
	ReplaceInt64Variables(&n1, &n2)
	fmt.Printf("After replace:\nn1: %d, n2: %d\n", n1, n2)
}

func ReplaceInt64Variables(n1, n2 *int64){ //2 3
	*n1 += *n2 //n1 = 5
	*n2 = *n1 - *n2 //n2 = 5 - 3 = 2
	*n1 -= *n2 //n1 = 5 - 2 = 3
}
//свойства умножения и сложения