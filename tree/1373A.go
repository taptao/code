package tree

import "math"

func maxSumBST(root *TreeNode) int {
	max := 0
	var bianli func(node *TreeNode) (int, int, int, bool)
	bianli =  func(node *TreeNode) (int,int,int,bool){
		if node == nil{
			return math.MaxInt,math.MinInt,0,true
		}
		lmin, lmax, lsum, lexits := bianli(node.Left)
		rmin, rmax, rsum, rexits := bianli(node.Right)

		if !lexits || !rexits{
			return 0,0,0,false
		}
		if lexits && rexits && lmax < node.Val && rmin > node.Val{
			if lsum + rsum + node.Val> max{
				max = lsum + rsum + node.Val
			}
		}


		if node.Right == nil && node.Left == nil{

			return node.Val, node.Val, node.Val, true
		}else if node.Right == nil{
			return lmin, node.Val,lsum+rsum+ node.Val, true

		}else if node.Left == nil{
			return node.Val,rmax,lsum+rsum+ node.Val, true
		}
		return lmin,rmax,lsum+rsum+ node.Val, true
	}
	_, lmax, lsum, lexits := bianli(root.Left)
	rmin, _, rsum, rexits := bianli(root.Right)


	if !lexits || !rexits{
		return max
	}
	if lexits && rexits && lmax < root.Val && rmin > root.Val{
		if lsum + rsum + root.Val> max{
			max = lsum + rsum + root.Val
		}
	}

	return max
}