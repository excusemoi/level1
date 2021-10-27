package task4

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Solve(){
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}
	jobs := make(chan int, n)
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL) //отслеживаем ctrl c

	go func() { //фоном слушаем канал прерывания, если из него придёт os.Interrupt
		<- interrupt //закрываем канал jobs => горутины завершаются
		close(jobs)
		time.Sleep(time.Second)
		os.Exit(0)
	}()

	for i := 0; i < n; i++ { //запускаем n воркеров
		go func(id int, j chan int) {
			for s := range j {
				fmt.Printf("Worker #%d -> %d\n", id, s)
				time.Sleep(time.Second)
			}
			fmt.Printf("\nWorker %d closed\n", id)
		}(i+1, jobs)
	}


	rand.Seed(time.Now().UnixNano())

		for {
			jobs <- rand.Intn(100)
			time.Sleep(time.Second*2)
		}

}