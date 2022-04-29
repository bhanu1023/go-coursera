/*
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const (
	allowed = 2
)

type Chops struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *Chops
}

func host(p Philo, ch1 chan int) {
	philospher := <-ch1
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("starting to eat %d \n", philospher)
	fmt.Printf("starting to eat %d \n", philospher)
	p.leftCS.Unlock()
	p.rightCS.Unlock()
}

func (p Philo) eat(philospher int, ch1 chan int) {
	for i := 0; i < 3; i++ {
		ch1 <- philospher
		host(p, ch1)
	}
	wg.Done()

}

func main() {
	channel := make(chan int, allowed)
	CSsticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CSsticks[i] = new(Chops)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSsticks[i], CSsticks[(i+1)%5]}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(i, channel)
	}
	wg.Wait()
}
