package dynamic

import "fmt"

func LongestPalindrome(s string) string {
	length := len(s)
	if length <= 0{
		return ""
	}
	dp := make([][]int, length)
	for i:=0;i<length;i++{
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}

	for i:=0;i<length;i++{
		for j:=0;j<=i;j++{

			if s[i] == s[j]{

				if (j+1) > (i-1){
					dp[j][i] = 1
				}else if dp[j+1][i-1] == 1 {
					dp[j][i] = 1
				}
			}
		}
	}

	max := 0
	start := 0
	end := 0
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			if dp[i][j] == 1 && max < (j-i+1){
				start = i
				end = j
				max = j - i + 1
			}
			fmt.Print(dp[i][j])
		}
		fmt.Println("")
	}

	return  s[start:end]
}
