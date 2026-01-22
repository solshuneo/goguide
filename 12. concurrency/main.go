package main

import (
	"fmt"
	"sync"
)

func main() {
	const reqs = 10000000
	balance := reqs
	var controller sync.WaitGroup
	controller.Add(reqs)
	for range reqs {
		go func() {
			balance -= 1
			controller.Add(-1)
		}()

	}
	// bank
	// user -> balance = 1000
	// user share account: 1000 users
	// each new user: -1 x 1000 (req)
	// left = 0
	controller.Wait()
	fmt.Printf("the left: %d - %d\n", reqs, balance)
}
