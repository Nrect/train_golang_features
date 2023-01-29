package easy

// https://leetcode.com/problems/running-sum-of-1d-array/

// RunningSum  my answer
func RunningSum(nums []int) []int {
	result := make([]int, len(nums))
	prevElement := 0

	for i, currentNum := range nums {
		currentNum += prevElement
		prevElement = currentNum
		result[i] = currentNum
	}

	return result
}

// other solution

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
