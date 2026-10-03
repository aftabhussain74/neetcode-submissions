func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	longest := 0

	for _, num := range nums {

		if !seen[num-1]{
			length :=1
			current:= num 

			for seen[current+1] {
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
