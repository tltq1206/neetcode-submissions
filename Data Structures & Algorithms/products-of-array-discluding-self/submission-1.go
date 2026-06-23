func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i< len(result); i++ {
		result[i] = 1
		for j := 0; j < len(result); j++ {
			if j != i {
			result[i] *= nums[j]
			}  
		}
	}
	return result
}
