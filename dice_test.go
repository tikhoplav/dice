package dice

import (
	"testing"
	"math/rand"
)

func BenchmarkRollRecursive(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		RollRecursive()
	}
}

func BenchmarkRoll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Roll()
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(42)
		rand.Intn(5)
	}
}