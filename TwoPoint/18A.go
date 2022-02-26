package TwoPoint

import "sort"

/**


 */
func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) <=3 {
		return result
	}

	numsLen := len(nums)
	sort.Ints(nums)
	first := 0
	sec := 1
	third := 2

	for ;first<numsLen-3;{
		for sec = first+1; nums[sec] != nums[first]; sec++{}
		for
	}
}
