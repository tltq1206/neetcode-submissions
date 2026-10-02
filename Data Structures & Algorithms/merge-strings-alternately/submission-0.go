func mergeAlternately(word1 string, word2 string) string {
	f,s := 0,0
	var res string
	for f < len(word1) && s < len(word2){
		res = res + string(word1[f]) + string(word2[s])
		f++
		s++
	}

	if len(word1) < len(word2){
		res = res + word2[s:]
	}

	if len(word2) < len(word1){
		res = res + word1[f:]
	}

	return res
}
