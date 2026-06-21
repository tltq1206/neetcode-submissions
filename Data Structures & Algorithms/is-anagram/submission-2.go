func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	countS := make(map[rune]int)
	countT := make(map[rune]int)
	for i, v := range s {
		countS[v]++
		countT[rune(t[i])]++
	}

	for i, v:= range countT {
		if countS[i] != v  {
			return false
		}
	}
	return true
}
