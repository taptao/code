package main

func removeElement(nums []int, val int) int {
	start := 0
	end := len(nums) - 1

	for start <= end{
		if nums[start] == val{
			nums[start] = nums[end]
			end --
			continue
		}
		start ++
	}
	return end
}