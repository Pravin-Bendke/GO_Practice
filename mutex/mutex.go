package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mu   sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.view += 1

}

func main() {
	var wg sync.WaitGroup

	myView := post{view: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myView.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myView.view)
}
