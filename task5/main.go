package task5

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Solve(){
	var (
		s 		= make(chan int, 1)
		start 	= time.Now()
		n 		= 0
		num 	= 0
	)

	fmt.Println("Enter program time")
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	for ;time.Since(start) < time.Duration(n) * time.Second; {
		s <- rand.Intn(10)
		num = <-s
		fmt.Printf("Received: '%d'\n", num)
	}
	fmt.Printf("Program lasted %v\n",
		time.Now().
		Sub(start).
		Round(time.
			Second))
}



