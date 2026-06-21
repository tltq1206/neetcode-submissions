func majorityElement(nums []int) int {
    occ := len(nums)/2
	occCount := make(map[int]int)
	var result int
	for _, val := range nums {
		occCount[val]++
		if occCount[val] > occ {
			result = val
		}
	}
	return result
}
