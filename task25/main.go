package task25

import (
	"fmt"
	"log"
	"time"
)

func Solve(){
	var n int

	fmt.Println("Enter seconds to sleep:")

	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Program lasted %v\n", CustomSleep1(time.Duration(n) * time.Second))


}

func CustomSleep(d time.Duration) time.Duration { //for
	start := time.Now()
	for ; time.Since(start) < d; {
	}
	return time.Now().Sub(start).Round(time.Second)
}

func CustomSleep1(d time.Duration) time.Duration { //for + select + time.After
	start := time.Now()
	for {
		select {
		case <-time.After(d):
			return time.Now().Sub(start).Round(time.Second)
		}
	}
}
//time.After(d Duration) вернёт Time по истечению d