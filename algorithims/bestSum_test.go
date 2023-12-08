package algorithms

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkBestSum(b *testing.B) {
	for i := 0; i < 1; i++ {
		r := &Result{
			CurrentNumber: 0,
			Value:         &[]int{},
			ValueState:    &[]int{},
			State:         0,
			Memoizer:      hashmap.New[int, *[]int](),
		}
		result := BestSum(7, &[]int{3, 4, 7}, r)
		// result := BestSum(8, &[]int{2, 3, 5}, r)
		// result := BestSum(8, &[]int{2, 3, 5}, r)
		// result := BestSum(7, &[]int{2, 4}, r)
		// result := BestSum(7, &[]int{5, 3, 4, 7}, r)
		// result := BestSum(7, &[]int{2, 3}, r)
		// result := BestSum(300, &[]int{7, 14}, r)
		fmt.Printf("%#v\n", result.Value)
	}
}
