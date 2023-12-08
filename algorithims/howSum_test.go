package algorithms

import (
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkHowSum(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := &Result{Value: &[]int{}, State: 0, Memoizer: hashmap.New[int, *[]int]()}
		_ = HowSum(300, &[]int{7, 14}, r)
		// result := HowSum(8, &[]int{2, 3, 5}, r)
		// result := HowSum(7, &[]int{2, 4}, r)
		// result := HowSum(7, &[]int{5, 3, 4, 7}, r)
		// result := HowSum(7, &[]int{2, 3}, r)
		// result := HowSum(300, &[]int{7, 14}, r)
		// fmt.Printf("%#v\n", result.Value)
	}

}
