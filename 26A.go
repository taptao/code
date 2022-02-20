package main

func removeDuplicates(nums []int) int {
	start:=0
	end := 0
	for end<len(nums){
		if end == len(nums) - 1{
			nums[start] = nums[end]
			start ++
			return start
		}
		if nums[end] != nums[end+1]{
			nums[start] = nums[end]
			start ++
		}
		end++
	}
	return start
}