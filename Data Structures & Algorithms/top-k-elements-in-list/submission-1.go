import "slices"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	type pair struct {
		num int
		count int
	}

	arr := []pair{}

	for num, count := range freq {
		arr = append(arr, pair{num, count})
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return b.count - a.count
	})

	result := make([]int, k)
	for i := 0; i< k; i++ {
		result[i] = arr[i].num
	} 
	return result
}
