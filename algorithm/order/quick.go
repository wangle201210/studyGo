package order

func quick(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return
}

func selec(nums []int) {
	min := 0
	for i := 0; i < len(nums)-1; i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
				break
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return
}
