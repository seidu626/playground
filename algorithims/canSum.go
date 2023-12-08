package algorithms

import "github.com/cornelk/hashmap"

func CanSum(targetSum int, numbers *[]int, memo *hashmap.Map[int, bool]) bool {
	fn, ok := memo.Get(targetSum)
	if ok {
		return fn
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range *numbers {
		remainder := targetSum - num
		result := CanSum(remainder, numbers, memo)
		memo.Set(remainder, result)
		if result {
			return true
		}
	}

	return false
}
