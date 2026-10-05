func lengthOfLongestSubstring(s string) int {
	//using sliding window
	window := make(map[byte]int)

	left:=0
	// right:=0

	maxLength:=0

	for right:=0;right<len(s);right++{
		ch:= s[right]

	// if duplicate character inside, then sliding window is invalid. need to move the left point

		if prevIndex, exists:= window[ch]; exists && prevIndex >= left 			{
			left = prevIndex +1
		}

		window[ch] = right

		length := right - left +1 

		if length > maxLength {
			maxLength = length
		}
		 
	}

	return maxLength




}
 