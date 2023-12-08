package algorithms

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	memo := hashmap.New[uint, uint]()
	result := Fib(50, memo)
	fmt.Println(result)
}

func BenchmarkFibNaive(b *testing.B) {
	result := FibNaive(1)
	fmt.Println(result)
}
