package algorithms

import (
	"fmt"
	"github.com/cornelk/hashmap"
)

func GridTraveler(m, n uint, memo *hashmap.Map[string, uint]) uint {
	key := fmt.Sprintf("%d,%d", m, n)
	reverseKey := fmt.Sprintf("%d,%d", n, m)
	fn, ok := memo.Get(key)
	if !ok {
		fn, ok = memo.Get(reverseKey)
	}
	if ok {
		return fn
	}

	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	memo.Set(key, GridTraveler(m-1, n, memo)+GridTraveler(m, n-1, memo))

	fn, ok = memo.Get(key)
	if !ok {
		fn, ok = memo.Get(reverseKey)
	}
	return fn
}

func GridTravelerWithoutReverseKey(m, n uint, memo *hashmap.Map[string, uint]) uint {
	key := fmt.Sprintf("%d,%d", m, n)
	fn, ok := memo.Get(key)
	if ok {
		return fn
	}

	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	memo.Set(key, GridTraveler(m-1, n, memo)+GridTraveler(m, n-1, memo))

	fn, ok = memo.Get(key)

	return fn
}
