package task22

import (
	"fmt"
	"math/big"
)

//использование пакета big

func Solve(){
	var a = big.NewInt(1048578)
	var b = big.NewInt(1048578)
	var a_ = big.NewInt(1048578)

	fmt.Println("Big multiplication:")
	fmt.Println(a.Mul(a_, b))
	fmt.Println("Big sum:")
	fmt.Println(a.Add(a_, b))
	fmt.Println("Big division:")
	fmt.Println(a.Div(a_, b))
	fmt.Println("Big subtraction:")
	fmt.Println(a.Sub(a_, b))

}