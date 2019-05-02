package modules

func ExtractNextToken(n int, runeSlice []rune) ([]rune, []rune) {
	var token []rune
	i := 0
	// get token
	for i = 0; len(token) < n && len(runeSlice) > i; i++ {
		if !IsIgnoredChar(runeSlice[i]) {
			token = append(token, runeSlice[i])
		}
	}
	// slide 1 rune
	runeSlice = runeSlice[1:]
	// slide until not ignored char
	for i, r := range runeSlice {
		if !IsIgnoredChar(r) {
			runeSlice = runeSlice[i:]
			break
		}
	}
	// ignore except n-gram
	if len(token) == n {
		return token, runeSlice
	} else {
		return nil, runeSlice
	}

}
