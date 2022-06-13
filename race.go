package main

import (
	"fmt"
	"sync"
)

func main() {
	scores := []int{0}
	
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("one")
		scores =append(scores, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("two")
		scores =append(scores, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("three")
		scores =append(scores, 3)
		wg.Done()
	}(wg)
	wg.Wait()
	fmt.Println(scores)
}