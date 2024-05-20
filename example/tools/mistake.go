package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func race1() {
	s := make([]int, 1)
	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
	}()
	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

func race2() {
	s := make([]int, 0, 1)
	go func() {
		s1 := append(s, 1)
		fmt.Println("s1:", s1)
	}()
	go func() {
		s2 := append(s, 2)
		fmt.Println("s2:", s2)
	}()
}

func race2Fix() {
	s := make([]int, 0, 1)
	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)
		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()
	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)
		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()
}

type Cache struct {
	mu       sync.RWMutex
	balances map[int64]float64
}

func (c *Cache) AddBalance(id int64, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances
	c.mu.RUnlock()
	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	averageBalance := sum / float64(len(balances))
	fmt.Println("averageBalance", averageBalance)
	return averageBalance
}

func race3() {
	cache := &Cache{
		mu: sync.RWMutex{},
		balances: map[int64]float64{
			1: 1,
			2: 2,
		},
	}
	go cache.AddBalance(3, 3)
	go cache.AverageBalance()
	time.Sleep(time.Second)
}

func race4() {
	wg := sync.WaitGroup{}
	var v uint64
	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1)
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v)
}

type Donation struct {
	cond    *sync.Cond
	balance int
}

func race5() {
	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			println("======")
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(3)
	go f(5)
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}

type Donation1 struct {
	mu      sync.RWMutex
	balance int
}

func race6() {
	donation := &Donation1{}
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	// go f(15)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
	time.Sleep(time.Second)
}

type Donation2 struct {
	balance int
	ch      chan int
}

func race7() {
	donation := &Donation2{ch: make(chan int)}
	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}
	go f(10)
	go f(15)
	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}
}

func race8() {

}
