package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	wg.Wait()
	fmt.Printf("Done")

}
