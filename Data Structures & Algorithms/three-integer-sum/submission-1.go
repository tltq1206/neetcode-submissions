func threeSum(nums []int) [][]int {
	res :=[][]int{}
	sort.Ints(nums)
	for i:=0 ; i<len(nums); i++{
		a := nums[i]
		if a > 0 {
			break
		}

		if i > 0 && a ==nums[i-1]{
			continue
		}

		l := i+1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res,[]int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l+1]{
					l++
				}
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}
	return res
}
