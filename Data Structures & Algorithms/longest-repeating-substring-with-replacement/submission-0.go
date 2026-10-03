func characterReplacement(s string, k int) int {
	seen := make(map[byte]int)

	left := 0
	maxFreq := 0
	result :=0

	for right:=0; right < len(s); right++{
		seen[s[right]]++

		maxFreq = max(maxFreq, seen[s[right]])

		windowSize := right - left +1

		// check if invalid
		if windowSize-maxFreq > k {
			seen[s[left]]--
			left++
		}

		result = max(result, right-left+1)
	}


	return result

}
