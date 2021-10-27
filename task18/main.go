package task18

import (
	"fmt"
	"log"
	"sync"
)

type Counter struct {
	i int
	m sync.Mutex
}

func (c *Counter) Inc() {
	c.m.Lock()//предоставляем доступ к счётчику только одной горутине
	c.i++
	c.m.Unlock()
}

func Solve(){
	var n int
	var wg sync.WaitGroup
	var c Counter = Counter{
		i: 0,
		m: sync.Mutex{},
	}
	fmt.Println("Enter n:")
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.i)
}