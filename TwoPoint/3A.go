package TwoPoint

// Given a string s, find the length of the longest substring without repeating characters.
//func main() {
//	fmt.Println(lengthOfLongestSubstring("abba"))
//
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	index := make(map[byte]int)
	start := 0
	end := 0
	maxlenth := 1
	index[s[0]] = 0

	for ;end<len(s);end++{
		if start == end{
			continue
		}
		value, err := index[s[end]]
		if err && value>=start{
			start = value + 1
		}
		index[s[end]] = end
		maxlenth = Max(maxlenth, end - start + 1)
	}
	return maxlenth

}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
