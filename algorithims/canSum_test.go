package algorithms

import (
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkCanSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := hashmap.New[int, bool]()
		_ = CanSum(300, &[]int{7, 14}, memo)
		// fmt.Println(result)
	}
}
