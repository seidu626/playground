package algorithms

import "github.com/cornelk/hashmap"

type Result struct {
	CurrentNumber int
	State         int
	Value         *[]int
	ValueState    *[]int
	Memoizer      *hashmap.Map[int, *[]int]
}

func HowSum(targetSum int, numbers *[]int, result *Result) *Result {
	fn, ok := result.Memoizer.Get(targetSum)
	if ok {
		result.Value = fn
		return result
	}

	if targetSum == 0 {
		result.State = 0
		return result
	}

	if targetSum < 0 {
		result.State = -1
		return result
	}

	for _, num := range *numbers {
		remainder := targetSum - num
		result = HowSum(remainder, numbers, result)
		if result.State == 0 {
			*result.Value = append(*result.Value, num)
			result.Memoizer.Set(remainder, result.Value)
			// return only one combination of numbers sum up to targetSum
			return result
		}
	}
	result.Memoizer.Set(targetSum, &[]int{})
	return result
}
