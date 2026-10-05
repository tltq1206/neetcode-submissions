func getConcatenation(nums []int) []int {
	var res []int
	res = append(res, nums...)
	res = append(res, nums...)
	return res
}
