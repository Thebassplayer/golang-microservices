package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0

	lock sync.Mutex

	atomicCounter atomicInt
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *atomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.value++
}
func (i *atomicInt) Decreae() {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.value--
}
func (i *atomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()

	fmt.Printf("Final counter: %d", counter)
	fmt.Printf("Final atomic counter value: %d", atomicCounter.Value())
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	atomicCounter.Increase()
	wg.Done()
}
