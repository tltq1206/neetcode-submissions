func sortColors(nums []int) {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] > nums[j] {
                temp := nums[i]
                nums[i] = nums[j]
                nums[j] = temp
            }
        }
    }
}