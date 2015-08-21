// +build ignore

package main

import "sync"

type LazyPrimes struct {
	once   sync.Once // Guards the primes slice.
	primes []int
}

func (p *LazyPrimes) init() {
	// Populate p.primes with prime numbers.
}

func (p *LazyPrimes) Primes() []int {
	p.once.Do(p.init)
	return p.primes
}
