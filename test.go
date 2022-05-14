//package main
//
//func main() {
//	a := make([]int, 10)
//	c := append(a, 1, 2, 3, 4, 5)
//	b := a[1:]
//	print(&a)
//	print(&b)
//	print(&c)
//
//
//}


package main

import "fmt"

func main() {
	// 数组的地址可以通过数组名来获取
	// 数组的第一个元素的地址就是数组的地址
	// 数组的各个元素的间隔是依据数组的类型决定的，比如int64间隔8个，int32间隔4个
	abc := make(map[string]string)
	i,j := abc["213"]
	fmt.Println(j)
	fmt.Println(i)
	intArr := []int {9,9,9}
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Printf("intArr的地址=%p int[0]地址=%p int[1]地址%p int[2]地址%p", &intArr, &intArr[0], &intArr[1], &intArr[2])

	c := append(intArr, 1,2,3)
	intArr = append(intArr, 1,2,4)
	fmt.Println()
	fmt.Printf("intArr的地址=%p int[0]地址=%p int[1]地址%p int[2]地址%p", &intArr, &intArr[0], &intArr[1], &intArr[2])

	fmt.Println()
	fmt.Printf("%+v", c)
	fmt.Printf("%+v", intArr)
	fmt.Printf("intArr的地址=%p int[0]地址=%p int[1]地址%p int[2]地址%p", &c, &c[0], &c[1], &c[2])
}