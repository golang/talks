package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"testing"
)

type Person struct {
	Name     string
	AgeYears int
	SSN      int
}

type byName []Person

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type byAge []Person

func (b byAge) Len() int           { return len(b) }
func (b byAge) Less(i, j int) bool { return b[i].AgeYears < b[j].AgeYears }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type bySSN []Person

func (b bySSN) Len() int           { return len(b) }
func (b bySSN) Less(i, j int) bool { return b[i].SSN < b[j].SSN }
func (b bySSN) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func BenchmarkSortSort(b *testing.B) {
	p := manyPeople()
	for i := 0; i < b.N; i++ {
		sort.Sort(byName(p))
		sort.Sort(byAge(p))
		sort.Sort(bySSN(p))
	}
}

func BenchmarkSortSlice(b *testing.B) {
	p := manyPeople()
	for i := 0; i < b.N; i++ {
		sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
		sort.Slice(p, func(i, j int) bool { return p[i].AgeYears < p[j].AgeYears })
		sort.Slice(p, func(i, j int) bool { return p[i].SSN < p[j].SSN })
	}
}

func manyPeople() []Person {
	n, err := strconv.Atoi(os.Getenv("N"))
	if err != nil {
		panic(err)
	}
	p := make([]Person, n)
	for i := range p {
		p[i].AgeYears = rand.Intn(100)
		p[i].SSN = rand.Intn(10000000000)
		p[i].Name = fmt.Sprintf("Mr or Ms %d", p[i].AgeYears)
	}
	return p
}
