func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)
	for k, v := range nums {
		indices[v] = k
	}
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
			res = append(res, indices[nums[l]], indices[nums[r]])
			break
		}
	}
	return res
}
