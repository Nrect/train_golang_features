package easy

// https://leetcode.com/problems/running-sum-of-1d-array/

// RunningSum  my answer
func RunningSum(nums []int) []int {
	prevElement := 0

	for i, currentNum := range nums {
		currentNum += prevElement
		prevElement = currentNum
		nums[i] = currentNum
	}

	return nums
}

// other solution better:
// 1) in-place operation - no additional memory allocated to store result
// 2) also no additional variable to store the ongoing sum

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
