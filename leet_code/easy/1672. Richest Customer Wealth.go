package easy

//https://leetcode.com/problems/richest-customer-wealth/
// testData := [][]int{{1, 2, 30}, {3, 2, 1}, {}}

// my solution

func MaximumWealth(accounts [][]int) int {
	maxAccountValue := 0

	for _, customer := range accounts {
		var customerBankSum int

		for _, bankValue := range customer {
			customerBankSum += bankValue
		}

		if customerBankSum > maxAccountValue {
			maxAccountValue = customerBankSum
		}
	}

	return maxAccountValue
}
