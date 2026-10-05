func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i<len(nums)-1;i++{
		a := nums[i]
		if a > target {
			break
		}
		if i > 0 && a == nums[i-1]{
			continue
		}
		for j := i+1; j < len(nums); j++ {
			if j > 1 && nums[j] == nums[j-1]{
				continue
			}
			l := j+1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else{
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1]{
						l++
					}
				}
			}
		}
	}
	return res
}
