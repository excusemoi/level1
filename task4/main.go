package task4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func Solve(){
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}
	jobs := make(chan string, n)
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt) //отслеживаем ctrl c

	for i := 0; i < n; i++ { //запускаем n воркеров
		go func(id int, j chan string) {
			for s := range j {
				fmt.Printf("Worker #%d -> %s\n", id, s)
				time.Sleep(time.Second)
			}
			fmt.Printf("\nWorker %d closed\n", id)
		}(i+1, jobs)
	}

	go func() { //фоном слушаем канал прерывания, если из него придёт os.Interrupt
		<- interrupt //закрываем канал jobs => горутины завершаются
		close(jobs)
		time.Sleep(time.Second)
		os.Exit(0)
	}()

	rd := bufio.NewReader(os.Stdin)

	for {
		s, _ := rd.ReadString('\n')
		jobs <- s
	}

}