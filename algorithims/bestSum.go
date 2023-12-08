package algorithms

import "fmt"

func BestSum(targetSum int, numbers *[]int, result *Result) *Result {
	if targetSum == 0 {
		result.State = 0
		return result
	}

	if targetSum < 0 {
		result.State = -1
		return result
	}

	for _, num := range *numbers {
		result.CurrentNumber = num
		remainder := targetSum - num
		if remainder < 0 {
			continue
		}
		fmt.Printf("TSum: %d Num: %d  CNum: %d\n", targetSum, num, remainder)
		result = BestSum(remainder, numbers, result)
		if result.State == 0 {
			*result.ValueState = append(*result.ValueState, num)
			fmt.Printf("State: 0 TSum: %d Num: %d  CNum: %d\n", targetSum, num, remainder)
			if len(*result.Value) > len(*result.ValueState) {
				*result.Value = *result.ValueState
				result.ValueState = new([]int)
			}
		}
	}

	return result
}
