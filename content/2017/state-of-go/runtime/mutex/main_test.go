// +build ignore,OMIT

package main

import (
	"fmt"
	"testing"
)

func benchFunc(b *testing.B, f func(int) map[int]int) {
	for n := 10; n <= 10000; n *= 10 {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}
}

func BenchmarkNarrowSection(b *testing.B) { benchFunc(b, countFactorsNarrowSection) }
func BenchmarkWideSection(b *testing.B)   { benchFunc(b, countFactorsWideSection) }
func BenchmarkSq(b *testing.B)            { benchFunc(b, countFactorsSeq) }
