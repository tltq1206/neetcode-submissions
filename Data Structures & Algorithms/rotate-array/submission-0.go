func rotate(nums []int, k int) {
	l := 0
	r := l + k
	for r < len(nums){
		temp := nums[l]
		nums[l] = nums[r]
		nums[r] = temp
		l++
		r = l + k
	}
}
