package TwoPoint

import "fmt"

func main() {
	a := "abcabca"
	b := "abcabcaaaaa"
	strStr(b,a)
}
func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}

	next := make([]int, len(needle))

	for i:=1;i<len(needle);i++{
		// abcabcd map[7] = 4
		//    abcabcd
		
	}
	fmt.Println(next)
	for i:=0;i<len(haystack);i++{
		j:=0
		for ;j<len(needle);j++{
			if i+j>=len(haystack){
				break
			}
			if needle[j] != haystack[i+j]{
				break
			}
		}
		if j == len(needle){
			return i
		}
	}
	return -1
}