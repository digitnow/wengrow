package bubble

import (
	"math/rand"
	"time"
	"testing"
)

func mockArrayOfInts(size int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(size)
}

func benchmarkBubble(i int, b *testing.B) {
	mockData := mockArrayOfInts(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
	    BubbleSort(mockData)
	}
}

func BenchmarkBubble10(b *testing.B) { benchmarkBubble(10, b) }
func BenchmarkBubble20(b *testing.B) { benchmarkBubble(20, b) }
func BenchmarkBubble40(b *testing.B) { benchmarkBubble(40, b) }
func BenchmarkBubble80(b *testing.B) { benchmarkBubble(80, b) }
func BenchmarkBubble100(b *testing.B) { benchmarkBubble(100, b) }
//func BenchmarkBubble1000000(b *testing.B) { benchmarkBubble(1000000, b) }