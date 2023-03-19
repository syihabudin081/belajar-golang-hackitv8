package rapih

import (
	"fmt"
	"sync"
	"time"
)

func RapihGoRoutine() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	interface1 := []interface{}{"bisal", "bisa2", "bisa3"}
	interface2 := []interface{}{"cobal", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go printData(interface1, &wg, &mu)
		go printData(interface2, &wg, &mu)
		wg.Wait()
	}
}

func printData(data []interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for j := len(data); j > 0; j-- {
		mu.Lock()
		fmt.Println(data, j)
		mu.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}