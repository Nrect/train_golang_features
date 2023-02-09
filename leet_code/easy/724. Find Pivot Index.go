package easy

//https://leetcode.com/problems/find-pivot-index/description/?envType=study-plan&id=level-1

// answer https://leetcode.com/problems/find-pivot-index/solutions/512992/python-go-java-js-c-o-n-sol-by-balance-scale-w-explanation/?envType=study-plan&id=level-1&orderBy=most_votes&languageTags=golang

// PivotIndex my solution
func PivotIndex(nums []int) int {

	leftWight := 0
	rightWeight := getNumSum(nums)

	for idx, currentWeight := range nums {
		rightWeight -= currentWeight

		if leftWight == rightWeight {
			return idx
		}

		leftWight += currentWeight
	}

	return -1
}

func getNumSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
