package services

import (
	"fmt"
	"sync"
)
func add( arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		*arr = append(*arr, i)
	}
}
func main() {
	var m []int
	var wg sync.WaitGroup
	wg.Add(2)
	go add(&m, &wg)
	go add(&m, &wg)

	wg.Wait()
	fmt.Println(m)
}
