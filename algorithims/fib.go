package algorithms

import "github.com/cornelk/hashmap"

func Fib(n uint, memo *hashmap.Map[uint, uint]) uint {
	fnv, ok := memo.Get(n)
	if ok {
		return fnv
	}
	if n <= 2 {
		return 1
	}
	memo.Set(n, Fib(n-1, memo)+Fib(n-2, memo))
	fnv, _ = memo.Get(n)
	return fnv
}

func FibNaive(n uint) uint {
	if n <= 2 {
		return 1
	}
	return FibNaive(n-1) + FibNaive(n-2)
}
