func twoSum(nums []int, target int) []int {
    sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	var res []int
	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			res = append(res, l, r)
			break
		}
	}
	return res
}
