package tree

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var backTrace func(sum int, ret []int, index int)
	backTrace = func(sum int,ret []int, index int) {
		if sum > target || len(ret) >= 150 {
			return
		}

		if sum == target {
			tmp := make([]int, len(ret))
			copy(tmp, ret)
			res = append(res, tmp)
			return
		}
		for i:=index;i<len(candidates);i++{
			num := candidates[i]
			c := append(ret, num)
			backTrace(sum+num,c, i)
		}
	}

	backTrace(0, []int{}, 0)

	return res
}