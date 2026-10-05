func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
	for i := 0; i<len(prefix); i++{
		for j:=0; j<len(strs); j++ {
			if i < len(strs[j]) && prefix[i] != strs[j][i]{
				prefix = prefix[:j]
				break
			}	
		}
	}
	return prefix
}
