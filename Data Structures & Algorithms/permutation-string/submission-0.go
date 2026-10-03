func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var need [26]int
	var window [26]int

	// counting all the freq of all chars in s1
	for i:=0;i<len(s1);i++ {
		need[s1[i]-'a']++
	}

	// making first window
	for i:=0;i< len(s1);i++ {
		window[s2[i]-'a']++
	}

	// checking first window
	if need == window {
		return true
	}

	// slide the window after the length of first
	for right:=len(s1); right < len(s2); right++ {
		window[s2[right]-'a']++

		left:= right -len(s1)
		window[s2[left]-'a']--

		if need == window {
			return true
		}
	}

	return false

}
