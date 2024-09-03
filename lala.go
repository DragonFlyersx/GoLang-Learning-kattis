package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork chan bool

type Philosopher struct {
	id        int
	leftFork  Fork
	rightFork Fork
	eatCount  int
}

func main11() {
	var wg sync.WaitGroup
	numPhilosophers := 5

	// Create forks
	forks := make([]Fork, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = make(Fork, 1)
		forks[i] <- true // Fork is initially available
	}

	// Create philosophers
	philosophers := make([]Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	// Start goroutines for each philosopher
	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go func(p Philosopher) {
			defer wg.Done()
			for p.eatCount < 3 {
				p.think()
				p.eat()
			}
		}(philosophers[i])
	}

	wg.Wait()
}

func (p *Philosopher) eat() {
	// Pick up left fork
	<-p.leftFork
	// Pick up right fork
	<-p.rightFork

	// Eat
	p.eatCount++
	fmt.Printf("Philosopher %d is eating. He has eaten %d times.\n", p.id, p.eatCount)
	time.Sleep(time.Second) // Simulate eating time

	// Put down right fork
	p.rightFork <- true
	// Put down left fork
	p.leftFork <- true
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking.\n", p.id)
	time.Sleep(time.Second) // Simulate thinking time
}
