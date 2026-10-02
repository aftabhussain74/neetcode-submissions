func twoSum(numbers []int, target int) []int {
	// make two pointer left and right. that way pointer won't overlap. and then in the loop if the sum of left and right is greater than target, the right--, if less, then left++.

	left := 0
	right := len(numbers) - 1

	for left < right {

		guess := numbers[left] + numbers[right]
		if guess == target {
			return []int{left + 1, right + 1}
		} 

		if guess < target {
			left++
		} else {
			right--
		}
	}

	 return []int{left + 1, right + 1}
}
