package TwoPoint

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3{
		return 0
	}
	one := 0
	minLen := math.MaxInt64
	res := 0
	sort.Ints(nums)
	for ;one<len(nums);one++{
		if one >= 1 && nums[one-1] == nums[one] {
			continue
		}
		two := one + 1
		third := len(nums) - 1
		sumNum := target - nums[one]
		for two<third{
			if two > one + 1 && nums[two] == nums[two-1]{
				two++
				continue
			}
			if minLen > Abs(sumNum - nums[two] - nums[third]){
				minLen = Abs(sumNum - nums[two] - nums[third])
				res = nums[one] + nums[two] + nums[third]
			}

			if sumNum > nums[two] + nums[third]{
				two ++
			}else if sumNum < nums[two] + nums[third]{
				third --
			}else{
				res = nums[one] + nums[two] + nums[third]
				return res
			}
		}
	}
	return res
}

func Abs(x int) int{
	if x>0{
		return x
	}
	return -x
}
