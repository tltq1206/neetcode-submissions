import "maps"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	listOccurrences := make([]map[rune]int, len(strs))

	for i, str := range strs {
		occurrences := make(map[rune]int)
		for _, v := range str {
			occurrences[v]++
		}
		listOccurrences[i] = occurrences
	}
	var result [][]string
	seen := make(map[int]bool)
	for i := 0; i < len(listOccurrences); i++ {
		if seen[i] { continue }
		currentGroup := []string{strs[i]}
		seen[i] = true
		for j := i + 1; j < len(listOccurrences); j++ {
			if !seen[j] {
				if maps.Equal(listOccurrences[i], listOccurrences[j]) {
					currentGroup = append(currentGroup, strs[j])
					seen[j] = true
				}
			}
		}
		result = append(result, currentGroup)
	}
	return result
}