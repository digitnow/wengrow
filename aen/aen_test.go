package aen

import (
	"math/rand"
	"time"
	"testing"
)

func mockArrayOfInts(size int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(size)
}

func benchmarkAen(i int, b *testing.B) {
	mockData := mockArrayOfInts(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
	    AverageOfEvenNumbers(mockData)
	}
}

func BenchmarkAen10(b *testing.B) { benchmarkAen(10, b) }
func BenchmarkAen100(b *testing.B) { benchmarkAen(100, b) }
func BenchmarkAen1000(b *testing.B) { benchmarkAen(1000, b) }
func BenchmarkAen10000(b *testing.B) { benchmarkAen(10000, b) }
func BenchmarkAen100000(b *testing.B) { benchmarkAen(100000, b) }
func BenchmarkAen1000000(b *testing.B) { benchmarkAen(1000000, b) }