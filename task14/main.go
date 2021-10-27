package task14

import (
	"fmt"
	"reflect"
)

func Solve(v interface{}) {
	fmt.Printf("Type of %v is %v\n", v, DetermineVariableType(v))
}

func DetermineVariableType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
