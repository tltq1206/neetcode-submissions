func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	countT := make(map[byte]int)
	countS := make(map[byte]int)
	for  i:= 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}

	for k, v := range countT {
		if v != countS[k]{
			return false
		}
	}
	return true
}
