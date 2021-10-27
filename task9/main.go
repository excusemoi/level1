package task9

import "fmt"

func Solve(){
	var in chan int = make(chan int)
	var out chan int = make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() { //заносив в канал in элементы arr
		for _, val := range arr {
			in <- val
		}
		close(in) //закрываем канал во избежание deadlock'a в дальнейшем
		//(ситуации, когда канал, из которого читают, никогда не закрывается)
	}()

	go func(){
		for n := range in {
			out <- 2 * n
		}
		close(out)
	}()

	for n := range out{
		fmt.Println(n)
	}

}

