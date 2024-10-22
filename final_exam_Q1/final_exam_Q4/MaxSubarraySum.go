package code

func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSum := arr[0]     // Initalize maxSum to the first element of the array
	currentSum := arr[0] // Current sum starts with the first element

	for i := 1; i < len(arr); i++ {
		// If currentSum becomes less than 0, reset it to the current element
		currentSum = max(arr[i], currentSum+arr[i])
		// Update maxSum if currentSum is larger
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// Helper function to get the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
