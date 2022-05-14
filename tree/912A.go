package tree

func sortArray(nums []int) []int {

	temp := make([]int, len(nums))

	var sort func(lo ,hi int)
	var merge func(lo, mid, hi int)

	sort = func(lo, hi int) {

		if lo >= hi{
			return
		}

		mid := lo + (hi - lo) / 2
		sort(lo, mid)
		sort(mid+1, hi)
		merge(lo, mid,hi)
	}

	merge = func(lo, mid, hi int) {
		for i:=lo;i<=hi;i++{
			temp[i] = nums[i]
		}
		i:=lo
		j:=mid+1
		for index:=lo;index<=hi;index++{
			if i == mid + 1{
				nums[index] = temp[j]
				j++
			}else if j == hi + 1{
				nums[index] = temp[i]
				i++
			}else if temp[i] < temp[j]{
				nums[index] = temp[i]
				i++
			}else{
				nums[index] = temp[j]
				j++
			}
		}
	}


	sort(0,len(nums)-1)

	return nums
}