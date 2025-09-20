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

	if q == p {
		return true
	}
	if q == nil || p == nil {
		return false
	}

	sp = append(sp, p)
	sq = append(sq, q)

	for len(sp) > 0 && len(sq) > 0 {
		cp := sp[len(sp)-1]
		sp = sp[:len(sp)-1]

		cq := sq[len(sq)-1]
		sq = sq[:len(sq)-1]

		if cp.Val != cq.Val {
			return false
		}

		if cp.Right != nil {
			sp = append(sp, cp.Right)
		}
		if cq.Right != nil {
			sq = append(sq, cq.Right)
		}
		if len(sq) != len(sp) {
			return false
		}

		if cp.Left != nil {
			sp = append(sp, cp.Left)
		}
		if cq.Left != nil {
			sq = append(sq, cq.Left)
		}
		if len(sq) != len(sp) {
			return false
		}
	}

	return true
}
