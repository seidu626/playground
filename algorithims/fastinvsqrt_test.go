package algorithms

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkFastInvSqrt32(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result := FastInvSqrt32(float32(i))
		fmt.Println(result)
	}
}

func BenchmarkFastInvSqrt64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result := FastInvSqrt64(float64(i))
		fmt.Println(result)
	}
}

func BenchmarkFastInvNormal(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result := math.Sqrt(float64(i))
		fmt.Println(result)
	}
}
