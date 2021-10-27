package task7

import (
	"fmt"
	"sync"
)

func Solve(){
	m := make(map[int]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			wg.Done()
			lock.Lock() //доступ к ресурсу m получает только одна горутина
			m[key] = key + 1
			lock.Unlock()
		}(i)
	}
	wg.Wait()
	for key, val := range m{
		fmt.Printf("Key: %d Value: %d\n", key, val)
	}

}