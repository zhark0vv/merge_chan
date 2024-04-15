package main

import (
	"fmt"
	"sync"
)

func main() {
	fillChan := func(inputCh chan int, numbers []int) {
		for _, n := range numbers {
			inputCh <- n
		}
		close(inputCh)
	}

	output := make(chan int)

	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go fillChan(first, []int{1, 3, 5})
	go fillChan(second, []int{3, 5, 7})
	go fillChan(third, []int{3, 5, 7})

	chans := []chan int{first, second, third}

	wg := sync.WaitGroup{}

	for _, forMerge := range chans {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for input := range ch {
				output <- input
			}
		}(forMerge)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for o := range output {
		fmt.Println(o)
	}
}
