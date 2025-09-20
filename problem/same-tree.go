package problem

// https://leetcode.com/problems/same-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	sp, sq := []*TreeNode{}, []*TreeNode{}

	sp = append(sp, p)
	sq = append(sq, q)

	for len(sp) > 0 && len(sq) > 0 {
		cp := sp[len(sp)-1]
		cq := sq[len(sq)-1]

		sp = sp[:len(sp)-1]
		sq = sq[:len(sq)-1]

		if cp == nil && cq == nil {
			continue
		}
		if cp == nil || cq == nil {
			return false
		}
		if cp.Val != cq.Val {
			return false
		}

		sp = append(sp, cp.Right)
		sp = append(sp, cp.Left)
		sq = append(sq, cq.Right)
		sq = append(sq, cq.Left)
	}

	return true
}
