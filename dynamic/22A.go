package dynamic


func generateParenthesis(n int) []string {
	if n <=0{
		return []string{}
	}
	var result []string
	create(n, 0, 0, "", &result)

	return result
}

func create(n ,rightNum ,leftNum int, rStr string, res *[]string){

	if rightNum == n && leftNum == n{
		*res = append(*res, rStr)
		return
	}


	if leftNum < n{
		create(n ,rightNum, leftNum + 1, rStr + "(", res)
	}

	if rightNum < n && rightNum < leftNum{
		create(n, rightNum+1, leftNum, rStr + ")", res)
	}
}

