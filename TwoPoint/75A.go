package TwoPoint
func sortColors(nums []int)  {
	before := 0
	after := len(nums)

	for i:=0;i<=after;{
		if nums[i]==0{
			nums[i] = nums[before]
			if i == before{
				i++
				before++
			}else{
				before++
			}
			nums[before] = 0
			continue
		}

		if nums[i]==2{
			nums[i] = nums[after]
			after --
			continue
		}

		i++
	}
}
