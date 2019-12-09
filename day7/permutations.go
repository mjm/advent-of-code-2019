package day7

func AllPermutations(nums []int) chan []int {
	out := make(chan []int)
	go func() {
		generatePermutations(len(nums), nums, out)
		close(out)
	}()
	return out
}

func generatePermutations(k int, nums []int, out chan []int) {
	if k == 1 {
		ret := make([]int, len(nums))
		copy(ret, nums)
		out <- ret
		return
	}

	generatePermutations(k-1, nums, out)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			nums[i], nums[k-1] = nums[k-1], nums[i]
		} else {
			nums[0], nums[k-1] = nums[k-1], nums[0]
		}
		generatePermutations(k-1, nums, out)
	}
}
