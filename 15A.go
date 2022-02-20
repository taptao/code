package main

import (
	"sort"
)
//
//func main() {
//	nums := []int{-1,0,1,2,-1,-4}
//	println(len(threeSum(nums)))
//}
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for sum:=0;sum<len(nums);sum++{
		if sum>=1 && (nums[sum-1] == nums[sum]){
			continue
		}
		end := len(nums) - 1
		numSum := nums[sum]
		for start:=sum ;start<end;{
			if start == sum{
				start ++
				continue
			}
			if end == sum{
				end --
				continue
			}
			if start>sum+1 && (nums[start-1] == nums[start]){
				start ++
				continue
			}

			if 0 > (nums[start] + nums[end]) + numSum{
				start ++
			}else if 0 < (nums[start] + nums[end]) + numSum{
				end --
			}else{
				 res = append(res, []int{nums[start], nums[end], nums[sum]})
				 start ++
			}
		}
	}
	return res
}