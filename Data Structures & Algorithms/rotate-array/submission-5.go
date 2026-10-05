func rotate(nums []int, k int) {
	pos := make([]int, len(nums))
	for i := 0; i< len(nums); i++{
		idx := (i+k)%len(nums)
		pos[idx] = nums[i]
	}

	for i := 0; i<len(nums); i++{
		nums[i] = pos[i]
	}

}
