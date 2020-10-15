package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func set(key string, value int) {
	m.Store(key, value)
}
func get(key string) int {
	if v, ok := m.Load(key); ok {
		return v.(int)
	}
	return 0
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 120; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("key: %q, value: %d\n", key, get(key))
		}(i)
	}
	wg.Wait()
}
