package algorithms

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"testing"
)

func BenchmarkGridTraveler(b *testing.B) {
	memo := hashmap.New[string, uint]()
	traveller := GridTraveler(18, 18, memo)
	fmt.Println(traveller)
}

func BenchmarkGridTravelerWithoutReverseKey(b *testing.B) {
	memo := hashmap.New[string, uint]()
	traveller := GridTravelerWithoutReverseKey(18, 18, memo)
	fmt.Println(traveller)
}
