package main

import (
	algorithms "github.com/seidu626/playground/algorithims"
	"math/rand"
	"testing"
)

var data  []int = rand.Perm(20000)

// BenchmarkBubbleSort https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.BubbleSort(&data)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.SelectionSort(&data)
	}
}