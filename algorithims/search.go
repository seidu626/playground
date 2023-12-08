package algorithms

// findDuplicates: Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice,
//return an array of all the integers that appears twice.
//You must write an algorithm that runs in O(n) time and uses only constant extra space.

// Note: subtracting 1 from each element in the array gives a valid index since the range is [1, n]
// 1: Walk through each element and obtain its position index (element - 1)
// 2: Negate all values found with the positional indexes
// 3: When we encounter a -value with a new positional index, this implies it is a duplicate

func FindDuplicates(nums []int) []int {
	var result []int
	for i := 0; i <= len(nums) - 1; i ++ {
		index := Abs(nums[i]) - 1
		if nums[index] < 0 {
			result = append(result, index + 1)
		}
		nums[index] = -1 * nums[index]
	}
	return  result
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}