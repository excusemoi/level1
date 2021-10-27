package task15

import (
	"fmt"
	"strconv"
)

var justString string

func SomeFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] отсутствие контроля за размером индекса
	justString = v[:len(v)-1] //так лучше
	fmt.Println(justString)
}

func createHugeString(n int64) string{
	return strconv.FormatInt(n, 2)
}


