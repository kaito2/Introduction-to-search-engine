package modules

func ExtractNextToken(start int, n int, runeSlice []rune) ([]rune, int) {
	var token []rune
	var i int
	// get token
	for i = start; len(token) < n && len(runeSlice) > i; i++ {
		if !IsIgnoredChar(runeSlice[i]) {
			token = append(token, runeSlice[i])
		}
	}
	// slide 1 rune
	nextStart := start + 1
	// slide until ignore rune
	for i = nextStart;i < len(runeSlice) ;i++ {
		if !IsIgnoredChar(runeSlice[i]) {
			nextStart = i
			break
		}
	}
	// ignore except n-gram
	if len(token) == n {
		return token, nextStart
	} else {
		return nil, nextStart
	}

}
