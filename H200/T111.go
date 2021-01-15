package H200

import (
	"math"
)

func MinDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	q := []*TreeNode{root}
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0;i<size;i++{
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil{
				return depth
			}
			if cur.Left != nil{
				q = append(q,cur.Left)
			}
			if cur.Right != nil{
				q = append(q,cur.Right)
			}
		}
		depth++
	}
	return depth
}

func minDepth(root *TreeNode) int  {
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil{
		return 1
	}

	minD := math.MaxInt32

	if root.Left != nil{
		minD = min(minD,minDepth(root.Left))
	}
	if root.Right != nil{
		minD = min(minD,minDepth(root.Right))
	}
	return minD + 1

}

func min(a,b int) int{
	if a > b{
		return b
	}
	return a
}
