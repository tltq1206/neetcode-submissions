func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m - 1
	r := n - 1
	i := m+n - 1
	for l >= 0 && r >=0 {
		if nums1[l] < nums2[r]{
			nums1[i] = nums2[r]
			r--
		} else {
			nums1[i] = nums1[l]
			l--
		}
		i--
	}

	if r >= 0 {
		for r >= 0 {
			nums1[i] = nums2[r]
			i--
			r--
		}	
	}
}
