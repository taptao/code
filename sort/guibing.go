package main

import "fmt"

func main() {
	nums := []int{3,4,1,5,7,2,4}

	for i := range nums {
		fmt.Print(nums[i])
	}
	fmt.Println()
	//sort(nums, 0, len(nums)-1)
	for i := range nums {
		fmt.Print(nums[i])
	}
}

func sort(nums []int, left int, right int) {
	if left == right{
		return
	}
	mid := (right + left) / 2
	sort(nums, left, mid)
	sort(nums, mid+1, right)

	var newNums []int
	i:=0
	j:=0
	for {
		if left+i > mid{
			break
		}
		if mid+1+j > right{
			break
		}
		if nums[left+i] > nums[mid+1+j]{
			newNums = append(newNums, nums[left+i])
			i++
		}else{
			newNums = append(newNums, nums[mid+1+j])
			j++
		}
	}
	for left+i <= mid{
		newNums = append(newNums, nums[left+i])
		i++
	}
	for mid+1+j <= right{
		newNums = append(newNums, nums[mid+1+j])
		j++
	}

	nums = append(append(nums[0:left], newNums...), nums[right:]...)
}

