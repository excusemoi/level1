package task6

import (
	"context"
	"fmt"
	"time"
)

func Solve(){
	fmt.Println("Method #1:")
	method1()
	fmt.Println("Method #2:")
	method2()
	fmt.Println("Method #3:")
	method3()
}

//специальный канал для выхода
func method1(){
	quit := make(chan bool)
	c := make(chan string)

	go func() {
		select {
		case <-quit:
			fmt.Println("Goroutine ends its work")
			return
		case str := <- c:
			fmt.Printf("Goroutine received %s\n", str)
		default:
			fmt.Println("Goroutine is fine")
		}
	}()
	quit <- true
}


//один канал, как для получения данных, так и для завершения горутины
func method2(){
	c := make(chan string)
	go func() {
		select {
		case str := <-c:
			switch str {
			case "Stop":
				fmt.Println("Goroutine ends its work")
				return
			default:
				fmt.Printf("Goroutine received %s\n", str)
			}
		default:
			fmt.Println("Goroutine is fine")
		}
	}()
	c <- "Stop"

}

//использование контекста
func method3(){
	forever := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine ends its work")
				forever <- struct{}{}
				return
			default:
				fmt.Println("Goroutine is fine")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	cancel()

	<-forever
}

