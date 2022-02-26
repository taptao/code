package TwoPoint


//Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
//
//Notice that you may not slant the container.
//
func maxArea(height []int) int {
	if len(height) <= 0{
		return 0
	}
	start := 0
	end := len(height) - 1
	maxV := Min(height[start], height[end]) * (end - start)

	for ;end!=start;{
		if height[end] > height[start]{
			start ++
		}else{
			end --
		}
		maxV = Max(maxV, Min(height[start], height[end]) * (end - start))
	}
	return maxV
}
//
//func Max(x int, y int) int{
//	if x>y{
//		return x
//	}
//	return y
//}

func Min(x int, y int) int{
	if x>y{
		return y
	}
	return x
}